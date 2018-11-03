package ru.aptu.xml.ast;

import ru.aptu.xml.ASTTree.TextSource;

public class AstInput extends AstNode {

	private final Ast_Node root;

	public AstInput(Ast_Node root, TextSource source, int line, int offset, int endoffset) {
		super(source, line, offset, endoffset);
		this.root = root;
	}

	public Ast_Node getRoot() {
		return root;
	}

	@Override
	public void accept(AstVisitor v) {
		if (!v.visit(this)) {
			return;
		}
		if (root != null) {
			root.accept(v);
		}
	}
}
