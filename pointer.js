let address = {
      city: null,
      province: null, 
      country: null
    };
    
    let address1 = {
      ...address,
      city: "Blitar",
      province: "Jawa Timur",
      country: "Indonesia"
    };

    let address2 = {
      ...address1,
      city: "Malang"
    }
    
    console.log(address1)
    console.log(address2)
    