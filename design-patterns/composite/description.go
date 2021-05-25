/*

Composite es un patrón de diseño estructural que te permite componer objetos
en estructuras de árbol y trabajar con esas estructuras como si fueran objetos individuales.

El uso del patrón Composite sólo tiene sentido cuando el modelo central de tu aplicación 
puede representarse en forma de árbol.

Por ejemplo, imagina que tienes dos tipos de objetos: Productos y Cajas. 
Una Caja puede contener varios Productos así como cierto número de Cajas más pequeñas. 
Estas Cajas pequeñas también pueden contener algunos Productos o incluso Cajas más pequeñas, 
y así sucesivamente.

Digamos que decides crear un sistema de pedidos que utiliza estas clases. 
Los pedidos pueden contener productos sencillos sin envolver, así como cajas llenas de productos... y otras cajas. 
¿Cómo determinarás el precio total de ese pedido?

Vamos a intentar imaginar el patrón Composite con un ejemplo de un sistema de archivos del sistema operativo. 
En el sistema de archivos hay dos tipos de objetos: archivos y carpetas. 
Hay casos en los que archivos y carpetas deben tratarse de la misma manera. 
Aquí es donde el patrón Composite resulta de utilidad.

Imagina que tienes que realizar una búsqueda de una palabra clave particular en tu sistema de archivos. 
Esta operación de búsqueda se aplica tanto a archivos como a carpetas. 
Para un archivo, buscará en los contenidos del archivo; 
para una carpeta, recorrerá todos los archivos de la carpeta para encontrar la palabra clave.
/*