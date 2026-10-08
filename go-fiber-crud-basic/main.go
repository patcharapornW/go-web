package main
 
import (
    "strconv"
 
    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v3/middleware/cors"
)
 
// สร้าง struct ของ Product เพื่อจำลองข้อมูลที่จะเก็บใน database
type Product struct {
    ID    int     `json:"id"`
    Name  string  `json:"name"`
    Price float32 `json:"price"`
}
 
// สร้างตัวแปรแบบ struct Product เพื่อเก็บข้อมูล โดยข้อมูลจะเป็น mock data
var products []Product
 
func main() {
    // สร้าง instance ของ fiber
    app := fiber.New()
 
    // ใส่ mock data ในตัวแปร products เพื่อมาใช้ทดสอบกับ API (crud) -----
    products = append(products, Product{ID: 1, Name: "iPhone 15 Pro Max", Price: 35000})
    products = append(products, Product{ID: 2, Name: "Oppo Reno 11 5G", Price: 12000.75})
    products = append(products, Product{ID: 3, Name: "Samsung Galaxy S23 Ultra", Price: 55000.50})
    //---------------------------------------------------------------
 
    // ใช้ middleware สำหรับ CORS กรณีมีการเรียกใช้งานจากคนละ Network
    app.Use(cors.New(cors.Config{
        AllowOrigins: []string{"*"},
        AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders: []string{"origin", "content-type", "accept", "authorization"},
    }))
 
    // สร้าง API สำหรับ CRUD ข้อมูล คือ กำหนด route และค่าต่างๆ ที่ต้อง response กลับไปให้ client-user
    // (R-Read) ดึงข้อมูลทั้งหมด
    app.Get("/api/products", getAllProduct)
 
    // (R-Read) ดึงข้อมูลแบบมีเงื่อนไข
    app.Get("/api/products/:id", getProductByID)
 
    // (C-Create) เพิ่มข้อมูล
    app.Post("/api/products", createProduct)
 
    // (U-Update) แก้ไขข้อมูล
    app.Put("/api/products/:id", updateProductByID)
 
    // Start Server
    if err := app.Listen(":5599"); err != nil {
        panic(err)
    }
}
 
// ---------------------------------------------
// ฟังก์ชัน handler สำหรับ CRUD ข้อมูล
// ฟังก์ชันดึงข้อมูลทั้งหมด
func getAllProduct(c fiber.Ctx) error {
    return c.JSON(products)
}
 
// ฟังก์ชันดึงข้อมูลแบบมีเงื่อนไข
func getProductByID(c fiber.Ctx) error {
    // เอาข้อมูล parameter ที่ส่งมาจาก client-user เก็บในตัวแปร
    id := c.Params("id")
 
    // validate ข้อมูลที่ส่งมา ว่าเป็นตัวเลขหรือไม่
    if _, err := strconv.Atoi(id); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": "Invalid product ID",
        })
    }
 
    // วนลูปตรวจสอบว่าข้อมูลมีไหม ถ้ามีก็ส่งข้อมูลนั้นกลับไป
    for _, productResponse := range products {
        if strconv.Itoa(productResponse.ID) == id {
            return c.Status(200).JSON(productResponse)
        }
    }
 
    // กรณีไม่มี
    return c.Status(404).JSON(fiber.Map{
        "message": "Product not found",
    })
}
 
// ฟังก์ชันเพิ่มข้อมูล
func createProduct(c fiber.Ctx) error {
    // สร้างตัวแปรเก็บข้อมูลที่ส่งมาจาก client-user
    product := new(Product)
 
    // อ่านข้อมูลจาก body ของ request และเก็บในตัวแปรที่สร้างไว้
    if err := c.Bind().Body(product); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": "Invalid product data",
        })
    }
 
    // เพิ่มข้อมูลเข้าไปในตัวแปร products ที่เราเริ่มต้นกำหนดไว้ด้านบน
    products = append(products, *product)
 
    // เพิ่มเสร็จแล้วให้ return กลับไปบอก client-user
    return c.Status(201).JSON(fiber.Map{
        "message": "Product created successfully",
    })
}
 
// ฟังก์ชันแก้ไขข้อมูลแบบมีเงื่อนไข
func updateProductByID(c fiber.Ctx) error {
    // เอาข้อมูล parameter ที่ส่งมาจาก client-user เก็บในตัวแปร
    id := c.Params("id")
 
    // validate ข้อมูลที่ส่งมา ว่าเป็นตัวเลขหรือไม่
    if _, err := strconv.Atoi(id); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": "Invalid product ID",
        })
    }
 
    // สร้างตัวแปรเก็บข้อมูลที่ส่งมาจาก client-user
    product := new(Product)
 
    // อ่านข้อมูลจาก body ของ request และเก็บในตัวแปรที่สร้างไว้
    if err := c.Bind().Body(product); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": "Invalid product data",
        })
    }
 
    // วนลูปตรวจสอบว่ามีข้อมูลที่ตรงกับ id ที่ส่งมาไหม ถ้ามีก็แก้ไขข้อมูล
    for i, productData := range products {
        if strconv.Itoa(productData.ID) == id {
            products[i].Name = product.Name
            products[i].Price = product.Price
 
            return c.Status(200).JSON(fiber.Map{
                "message": "Product updated successfully",
            })
        }
    }
	// กรณีไม่มี หรือ หา id ไม่เจอ
    return c.Status(404).JSON(fiber.Map{
        "message": "Product not found",
    })
}


 // ฟังก์ชันการลบข้อมูลแบบมีเงื่อนไข
func deleteProductByID(c fiber.Ctx) error {
    // เอาข้อมูล parameter ที่ส่งมาจาก client-user เก็บในตัวแปร
    id := c.Params("id")

	// validate ข้อมูลที่ส่งมา ว่าเป็นตัวเลขหรือไม่
	if _, err := strconv.Atoi(id); err != nil {
        return c.Status(400).JSON(fiber.Map{
            "message": "Invalid product ID",
        })
	}

	 // วนลูปตรวจสอบว่ามีข้อมูลที่ตรงกับ id ที่ส่งมาไหม ถ้ามีก็ลบข้อมูล
    for i, productData := range products {
        if strconv.Itoa(productData.ID) == id {
            products = append(products[:i], products[i+1:]...)
 
            return c.Status(200).JSON(fiber.Map{
                "message": "Product deleted successfully",
            })
        }
    }

  // กรณีไม่มี หรือ หา id ไม่เจอ
    return c.Status(404).JSON(fiber.Map{
        "message": "Product not found",
    })
}


