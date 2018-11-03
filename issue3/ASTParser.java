package ru.aptu.xml;

import java.io.IOException;
import java.text.MessageFormat;
import ru.aptu.xml.ASTLexer.ErrorReporter;
import ru.aptu.xml.ASTLexer.Span;
import ru.aptu.xml.ASTLexer.Tokens;
import ru.aptu.xml.ast.AstInput;
import ru.aptu.xml.ast.Ast_Node;

public class ASTParser {

	public static class ParseException extends Exception {
		private static final long serialVersionUID = 1L;

		public ParseException() {
		}
	}

	private final ErrorReporter reporter;

	public ASTParser(ErrorReporter reporter) {
		this.reporter = reporter;
	}

	private static final boolean DEBUG_SYNTAX = false;
	private static final int[] tmAction = ASTLexer.unpack_int(7,
		"\uffff\uffff\uffff\uffff\0\0\uffff\uffff\1\0\uffff\uffff\ufffe\uffff");

	private static final int[] tmGoto = ASTLexer.unpack_int(8,
		"\0\0\2\0\4\0\6\0\10\0\10\0\12\0\14\0");

	private static final int[] tmFromTo = ASTLexer.unpack_int(12,
		"\5\0\6\0\1\0\3\0\0\0\1\0\3\0\4\0\0\0\5\0\0\0\2\0");

	private static final int[] tmRuleLen = ASTLexer.unpack_int(2,
		"\1\0\3\0");

	private static final int[] tmRuleSymbol = ASTLexer.unpack_int(2,
		"\5\0\6\0");

	protected static final String[] tmSymbolNames = new String[] {
		"eoi",
		"identifier",
		"openChar",
		"closeChar",
		"_skip",
		"input",
		"node",
	};

	public interface Nonterminals extends Tokens {
		// non-terminals
		int input = 5;
		int node = 6;
	}

	/**
	 * -3-n   Lookahead (state id)
	 * -2     Error
	 * -1     Shift
	 * 0..n   Reduce (rule index)
	 */
	protected static int tmAction(int state, int symbol) {
		return tmAction[state];
	}

	protected static int gotoState(int state, int symbol) {
		int min = tmGoto[symbol], max = tmGoto[symbol + 1];
		int i, e;

		while (min < max) {
			e = (min + max) >> 2 << 1;
			i = tmFromTo[e];
			if (i == state) {
				return tmFromTo[e+1];
			} else if (i < state) {
				min = e + 2;
			} else {
				max = e;
			}
		}
		return -1;
	}

	protected int tmHead;
	protected Span[] tmStack;
	protected Span tmNext;
	protected ASTLexer tmLexer;

	public AstInput parse(ASTLexer lexer) throws IOException, ParseException {

		tmLexer = lexer;
		tmStack = new Span[1024];
		tmHead = 0;

		tmStack[0] = new Span();
		tmStack[0].state = 0;
		tmNext = tmLexer.next();

		while (tmStack[tmHead].state != 6) {
			int action = tmAction(tmStack[tmHead].state, tmNext.symbol);

			if (action >= 0) {
				reduce(action);
			} else if (action == -1) {
				shift();
			}

			if (action == -2 || tmStack[tmHead].state == -1) {
				break;
			}
		}

		if (tmStack[tmHead].state != 6) {
			reporter.error(MessageFormat.format("syntax error before line {0}",
								tmLexer.getTokenLine()), tmNext.line, tmNext.offset, tmNext.endoffset);
			throw new ParseException();
		}
		return (AstInput)tmStack[tmHead - 1].value;
	}

	protected void shift() throws IOException {
		tmStack[++tmHead] = tmNext;
		tmStack[tmHead].state = gotoState(tmStack[tmHead - 1].state, tmNext.symbol);
		if (DEBUG_SYNTAX) {
			System.out.println(MessageFormat.format("shift: {0} ({1})", tmSymbolNames[tmNext.symbol], tmLexer.tokenText()));
		}
		if (tmStack[tmHead].state != -1 && tmNext.symbol != 0) {
			tmNext = tmLexer.next();
		}
	}

	protected void reduce(int rule) {
		Span left = new Span();
		left.value = (tmRuleLen[rule] != 0) ? tmStack[tmHead + 1 - tmRuleLen[rule]].value : null;
		left.symbol = tmRuleSymbol[rule];
		left.state = 0;
		if (DEBUG_SYNTAX) {
			System.out.println("reduce to " + tmSymbolNames[tmRuleSymbol[rule]]);
		}
		Span startsym = (tmRuleLen[rule] != 0) ? tmStack[tmHead + 1 - tmRuleLen[rule]] : tmNext;
		left.line = startsym.line;
		left.offset = startsym.offset;
		left.endoffset = (tmRuleLen[rule] != 0) ? tmStack[tmHead].endoffset : tmNext.offset;
		applyRule(left, rule, tmRuleLen[rule]);
		for (int e = tmRuleLen[rule]; e > 0; e--) {
			tmStack[tmHead--] = null;
		}
		tmStack[++tmHead] = left;
		tmStack[tmHead].state = gotoState(tmStack[tmHead - 1].state, left.symbol);
	}

	@SuppressWarnings("unchecked")
	protected void applyRule(Span tmLeft, int ruleIndex, int ruleLength) {
		switch (ruleIndex) {
			case 0:  // input : node
				tmLeft.value = new AstInput(
						((Ast_Node)tmStack[tmHead].value) /* root */,
						null /* input */, tmStack[tmHead].line, tmStack[tmHead].offset, tmStack[tmHead].endoffset);
				break;
			case 1:  // node : openChar identifier closeChar
				tmLeft.value = new Ast_Node(
						null /* input */, tmStack[tmHead - 2].line, tmStack[tmHead - 2].offset, tmStack[tmHead].endoffset);
				break;
		}
	}
}
