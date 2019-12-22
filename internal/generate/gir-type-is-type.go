package generate

func (t *Type) isAlias() bool {
	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Aliases.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Aliases.byName(t.Name)
	return found
}

func (t *Type) isBitfield() bool {
	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Bitfields.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Bitfields.byName(t.Name)
	return found
}

func (t *Type) isEnumeration() bool {
	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Enumerations.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Enumerations.byName(t.Name)
	return found
}

func (t *Type) isClass() bool {
	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Classes.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Classes.byName(t.Name)
	return found
}

func (t *Type) isRecord() bool {
	//if t.isClass() {
	//	return true
	//}

	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Records.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Records.byName(t.Name)
	return found
}

func (t *Type) isInterface() bool {
	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Interfaces.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Interfaces.byName(t.Name)
	return found
}

func (t *Type) isCallback() bool {
	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Callbacks.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Callbacks.byName(t.Name)
	return found
}

func (t *Type) isUnion() bool {
	if t.isQualifiedName() {
		_, found := t.foreignNamespace.Unions.byName(t.foreignName)
		return found
	}

	_, found := t.namespace.Unions.byName(t.Name)
	return found
}

func (t *Type) isVaList() bool {
	return t.Name == "va_list"
}