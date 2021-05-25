/*
	El patrón Builder se utiliza cuandl producto deseado es complejo y se requiren varios pasos para completarlo
	En este caso es más sencillo emplear varios métodos de construcción que un único constructor enorme.
	El problema con este proceso es que un producto parcialmente creado puede estar expuesto al cliente.
	El patrón Builder mantiene privado el producto hasta que esté totalmente creado

	En este código, veamos distintos tipos de casas (igloo, normalHouse) que son construidos por
	iglooBuilder y normalBuilder. Cada tipo de casa tiene los mismos pasos de construcción. 
	La estructura directora opcional ayuda a organizar el proceso de creación.
*/

