package ru.aptu.xml.ast;

import ru.aptu.xml.ASTTree.TextSource;

public class Ast_Node extends AstNode {

	public Ast_Node(TextSource source, int line, int offset, int endoffset) {
		super(source, line, offset, endoffset);
	}

	@Override
	public void accept(AstVisitor v) {
		v.visit(this);
	}
}
