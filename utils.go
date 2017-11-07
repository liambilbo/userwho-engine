package userwho_engine

func (addresses *Addresses) addAddress(address *Address,addressType AddressType) {
	paddrr:=PersonAddress{}
	paddrr.Address=*address
	paddrr.Type=addressType
	addresses.Addresses=append(addresses.Addresses,paddrr)
}

func (addresses *Addresses) Address(addressType AddressType) []PersonAddress{
	var res []PersonAddress
	for _,paddr := range addresses.Addresses {
		if paddr.Type ==addressType {
			res=append(res,paddr)
		}
	}

	return res
}

func (documents *Documents) addDocument(document *Document) {
	if documents.Documents==nil {
		documents.Documents=make(map[DocumentType]Document)
	}
}