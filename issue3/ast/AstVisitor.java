package ru.aptu.xml.ast;

public abstract class AstVisitor {

	protected boolean visit(AstInput n) {
		return true;
	}

	protected boolean visit(Ast_Node n) {
		return true;
	}
}
