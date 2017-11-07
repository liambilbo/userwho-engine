package userwho_engine


//Add an address to the person
func (addresses *Addresses) addAddress(address Address,addressType AddressType) {
	paddrr:=PersonAddress{}
	paddrr.Address=address
	paddrr.Type=addressType
	addresses.Addresses=append(addresses.Addresses,paddrr)
}

//Get the PersonAdress of a type
func (addresses *Addresses) Address(addressType AddressType) []PersonAddress{
	var res []PersonAddress
	for _,paddr := range addresses.Addresses {
		if paddr.Type ==addressType {
			res=append(res,paddr)
		}
	}

	return res
}

//Add a document to a person
func (documents *Documents) addDocument(document Document) {
	if documents.Documents==nil {
		documents.Documents=make(map[DocumentType]Document)
	}
	documents.Documents[document.Type]=document
}

//Get a document to a person
func (documents *Documents) Document(documentType DocumentType) (document Document) {
	document , _ = documents.Documents[documentType]
	return
}

