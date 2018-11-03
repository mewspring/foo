package ru.aptu.xml.ast;

import ru.aptu.xml.ASTTree.TextSource;

public interface IAstNode {
	String getLocation();
	int getLine();
	int getOffset();
	int getEndoffset();
	TextSource getSource();
	String getResourceName();
	String getText();
	void accept(AstVisitor v);
}
