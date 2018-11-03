package ru.aptu.xml;

import java.io.CharArrayReader;
import java.io.IOException;
import java.io.Reader;
import java.io.StringReader;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import ru.aptu.xml.ASTLexer.ErrorReporter;
import ru.aptu.xml.ASTParser.ParseException;
import ru.aptu.xml.ast.AstInput;

public class ASTTree<T> {

	private final TextSource source;
	private final T root;
	private final List<ASTProblem> errors;

	public ASTTree(TextSource source, T root, List<ASTProblem> errors) {
		this.source = source;
		this.root = root;
		this.errors = errors;
	}

	public TextSource getSource() {
		return source;
	}

	public T getRoot() {
		return root;
	}

	public List<ASTProblem> getErrors() {
		return errors;
	}

	public boolean hasErrors() {
		return errors.size() > 0;
	}


	public static ASTTree<AstInput> parse(TextSource source) {
		final List<ASTProblem> list = new ArrayList<>();
		ErrorReporter reporter = (message, line, offset, endoffset) ->
				list.add(new ASTProblem(KIND_ERROR, message, line, offset, endoffset, null));

		try {
			ASTLexer lexer = new ASTLexer(source.getContents(), reporter);
			lexer.setLine(source.getInitialLine());

			ASTParser parser = new ASTParser(reporter);
			AstInput result = parser.parse(lexer);

			return new ASTTree<>(source, result, list);
		} catch (ParseException ex) {
			/* not parsed */
		} catch (IOException ex) {
			list.add(new ASTProblem(KIND_FATAL, "I/O problem: " + ex.getMessage(), 0, 0, 0, ex));
		}
		return new ASTTree<>(source, null, list);
	}


	public static final int KIND_FATAL = 0;
	public static final int KIND_ERROR = 1;
	public static final int KIND_WARN = 2;

	public static final String PARSER_SOURCE = "parser";

	public static class ASTProblem extends Exception {
		private static final long serialVersionUID = 1L;

		private final int kind;
		private final int line;
		private final int offset;
		private final int endoffset;

		public ASTProblem(int kind, String message, int line, int offset, int endoffset, Throwable cause) {
			super(message, cause);
			this.kind = kind;
			this.line = line;
			this.offset = offset;
			this.endoffset = endoffset;
		}

		public int getKind() {
			return kind;
		}

		public int getLine() {
			return line;
		}

		public int getOffset() {
			return offset;
		}

		public int getEndoffset() {
			return endoffset;
		}

		public String getSource() {
			return PARSER_SOURCE;
		}
	}

	public static class TextSource {

		private final String file;
		private final int initialLine;
		private final CharSequence contents;
		private int[] lineoffset;

		public TextSource(String file, CharSequence contents, int initialLine) {
			this.file = file;
			this.initialLine = initialLine;
			this.contents = contents;
		}

		public String getFile() {
			return file;
		}

		public int getInitialLine() {
			return initialLine;
		}

		public Reader getStream() {
			if (contents instanceof String) {
				return new StringReader((String) contents);
			} else {
				return new CharArrayReader(contents.toString().toCharArray());
			}
		}

		public String getLocation(int offset) {
			return file + "," + lineForOffset(offset);
		}

		public String getText(int start, int end) {
			if (start < 0 || start > end || end > contents.length()) {
				return "";
			}
			return contents.subSequence(start, end).toString();
		}

		public int lineForOffset(int offset) {
			if (lineoffset == null) {
				lineoffset = getLineOffsets(contents);
			}
			int line = Arrays.binarySearch(lineoffset, offset);
			return initialLine + (line >= 0 ? line : -line - 2);
		}

		public int columnForOffset(int offset) {
			if (lineoffset == null) {
				lineoffset = getLineOffsets(contents);
			}
			int line = Arrays.binarySearch(lineoffset, offset);
			return offset >= 0 ? offset - lineoffset[line >= 0 ? line : -line - 2] : 0;
		}

		public CharSequence getContents() {
			return contents;
		}
	}

	private static int[] getLineOffsets(CharSequence contents) {
		int size = 1;
		int len = contents.length();
		for (int i = 0; i < len; i++) {
			if (contents.charAt(i) == '\n') {
				size++;
			} else if (contents.charAt(i) == '\r') {
				if (i + 1 < len && contents.charAt(i + 1) == '\n') {
					i++;
				}
				size++;
			}
		}
		int[] result = new int[size];
		result[0] = 0;
		int e = 1;
		for (int i = 0; i < len; i++) {
			if (contents.charAt(i) == '\n') {
				result[e++] = i + 1;
			} else if (contents.charAt(i) == '\r') {
				if (i + 1 < len && contents.charAt(i + 1) == '\n') {
					i++;
				}
				result[e++] = i + 1;
			}
		}
		if (e != size) {
			throw new IllegalStateException();
		}
		return result;
	}
}
