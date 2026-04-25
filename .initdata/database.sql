-- ==========================================
-- 1. User
-- ==========================================
CREATE TABLE users_table (
    UserID SERIAL PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Email VARCHAR(150) UNIQUE NOT NULL,
    Password VARCHAR(255) NOT NULL,
    PhoneNumber VARCHAR(10),
    AccountType VARCHAR(10) NOT NULL CHECK (AccountType IN ('Buyer', 'Seller', 'Admin')),
    Active VARCHAR(10) NOT NULL CHECK (Active IN ('active','ban'))
);

INSERT INTO users_table(Name,Email,Password,AccountType,Active)
VALUES
('TuanAmin','TuanAdmin@gmail.com','123','Admin','active'),
('Tuanseller','TuanSseller@gmail.com','123','Seller','ban');

-- ==========================================
-- 2. Shop
-- ==========================================
CREATE TABLE Shop (
    ShopID SERIAL PRIMARY KEY,
    ShopName VARCHAR(150) NOT NULL,
    Address VARCHAR(255) NOT NULL,
    Description TEXT,
    Logo VARCHAR(255),
    SellerID INT UNIQUE NOT NULL,
    FOREIGN KEY (SellerID) REFERENCES users_table(UserID)
);

INSERT INTO Shop(ShopName,Address,Description,SellerID)
VALUES
('GrabFood','120 Yen Lan','Den tu Singapore',1),
('ShopeFood','36 Dong hoa','Den tu Malaysia',2);
-- ==========================================
-- 3. Product
-- ==========================================
CREATE TABLE Product (
    ProductID SERIAL PRIMARY KEY,
    ProductName VARCHAR(150) NOT NULL,
    Price NUMERIC(12,2) NOT NULL,
    Description TEXT,
    image_url TEXT,
    StockQuantity INT NOT NULL DEFAULT 0,
    ShopID INT NOT NULL,

    FOREIGN KEY (ShopID) REFERENCES Shop(ShopID)
);
INSERT INTO product(ProductName,Price,image_url,StockQuantity,ShopID)
VALUES
('Snack vị bò nướng Texas Lays 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/226184/bhx/snack-khoai-tay-vi-than-bo-nuong-texas-lays-wavy-goi-56g_202508211052036366.jpg',8,1),
('Snack vị tự nhiên Lays 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/193606/bhx/snack-poca-khoai-tay-tu-nhien-goi-52g_202510150939206616.jpg',5,1),
('Snack vị sườn BBQ Lays 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/193603/bhx/snack-poca-khoai-tay-suon-nuong-52g_202509241622313039.jpg',6,1),
('Snack vị tảo biển Lays 50g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/193602/bhx/thumb-193602_202411261120181132.jpg',8,1),
('Snack vị tảo biển OStar 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/79812/bhx/snack-ostar-rong-bien-48g_202511181412004013.jpg',10,1),
('Snack vị kim chi OStar 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/79804/bhx/snack-khoai-tay-ostar-vi-kim-chi-han-quoc-48g_202511181404222381.jpg',3,1),
('Snack vị bít tết Swing 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/79773/bhx/snack-swing-nysteak-48g_202511251647001007.jpg',5,1),
('Snack vị bò lúc lắc Poca 65g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/283009/bhx/thumb-283009_202411261130399401.jpg',6,1),
('Snack vị tảo biển Orion 56g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/305485/bhx/thumb-305485_202411261132325426.jpg',13,1),
('Snack vị bò Wagyu Lays Max 42g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/313108/bhx/thumb-313108_202411261134209919.jpg',8,1),
('Snack sốt Mayo kiểu Hàn OStar gói 50g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/327381/bhx/thumb-327381-1_202411261351494616.jpg',2,1),
('Snack tôm nướng ngũ vị Lays 42g', 22000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/341877/bhx/snack-khoai-tay-tom-hum-nuong-ngu-vi-lays-stax-lon-42g_202508041450197713.jpg',13,1),
('Snack vị tự nhiên Lays Stax 42g', 22000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/341873/bhx/snack-khoai-tay-vi-tu-nhien-lays-stax-lon-42g_202508041439104475.jpg',15,1),
('Snack vị kem chua & hành Lays 42g', 22000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/341658/bhx/snack-khoai-tay-vi-kem-chua-hanh-lays-stax-lon-42g_202508041427115966.jpg',1,1),
('Snack vị phô mai Cheddar Lays 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/339156/bhx/snack-khoai-tay-vi-pho-mai-cheddar-lays-wavy-goi-54g_202510150943599559.jpg',17,1),
('Snack vị tôm càng đút phô mai Lays 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/339146/bhx/snack-khoai-tay-vi-tom-cang-dut-lo-pho-mai-tan-chay-lays-goi-54g_202510150952082071.jpg',14,1),
('Snack mũ pháp sư vị phô mai Poca 58g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/339143/bhx/snack-hinh-mu-phap-su-vi-pho-mai-poca-goi-58g_202506041018517919.jpg',18,1),
('Snack mực lăn muối ớt Poca 60g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/339136/bhx/snack-muc-lan-muoi-ot-poca-goi-60g_202506031456283223.jpg',2,1),
('Snack hải sản vị cay ngọt Bento 22g', 29000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/338657/bhx/snack-hai-san-cay-ngot-bento-goi-22g_202505270844320036.jpg',6,1),
('Snack hải sản vị cay nồng Bento 22g', 29000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/338656/bhx/snack-hai-san-vi-cay-nong-bento-goi-22g_202505270857071401.jpg',4,1),
('Snack bắp vị phô mai cay Koikeya Karamucho 60g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/338644/bhx/snack-bap-vi-pho-mai-cay-koikeya-karamucho-goi-60g_202505261441423139.jpg',15,1),
('Snack bắp Koikeya Polinky 54g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/338642/bhx/snack-bap-koikeya-polinky-goi-54g_202505261422442891.jpg',3,1),
('Bánh snack vị phô mai Nabati 80g', 24500, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/338613/bhx/banh-snack-vi-pho-mai-nabati-siip-hop-80g-4g-x-20-goi_202505260954160507.jpg',13,1),
('Bánh snack vị bắp rang Nabati 80g', 24500, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/338611/bhx/banh-snack-vi-bap-rang-nabati-siip-hop-80g-4g-x-20-goi_202505260943443524.jpg',7,1),
('Snack bắp vị socola Mplus 260g', 36000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/334868/bhx/snack-bap-vi-socola-mplus-chocolate-popcorn-chai-260g_202503281625537261.jpg',17,1),
('Snack hải sản Namprik Thái Bento 22g', 29000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/333971/bhx/thumb-333577-1_202501151003246540.jpg',10,1),
('Snack tôm xốt tương mật ong Oishi 70g', 12300, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/333564/bhx/thumb-333564-1_202412251002083319.jpg',10,1),
('Snack đậu xanh nước dừa Oishi 65g', 12300, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/333559/bhx/snack-dau-xanh-nuoc-dua-oishi-goi-70g_202506131528394847.jpg',15,1),
('Snack mực vị tương ớt Sriracha Talaethong 100g', 32000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/332876/bhx/snack-muc-vi-tuong-ot-sriracha-talaethong-goi-100g_202507281349450007.jpg',13,1),
('Snack bắp vị kem Mplus 260g', 36000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/327619/bhx/snack-bap-vi-kem-mplus-creamy-popcorn-chai-260g_202503281629438389.jpg',12,1),
('Snack bắp vị caramen Mplus 260g', 36000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/327616/bhx/thumb-327616_202411261403075655.jpg',8,1),
('Snack tôm vị cay nồng Oishi', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/305937/bhx/snack-tom-vi-cay-nong-oishi-goi-70g_202505270833459443.jpg',3,1),
('Snack socola vị hazelnut Oishi Pillows 85g', 12300, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/305932/bhx/thumb-305932_202411261132580586.jpg',11,1),
('Snack kem Tiramisu Akiko Oishi 140g', 24000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/305735/bhx/thumb-305735_202411261132415506.jpg',8,1),
('Snack kem Mixed Berries Akiko Oishi 140g', 24000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/305733/bhx/thumb-305733_202411261132362500.jpg',19,1),
('Snack mực rong biển Talaethong 100g', 31000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/282757/bhx/snack-muc-vi-rong-bien-talaethong-goi-100g_202507281356252761.jpg',7,1),
('Snack mực Talaethong 100g', 31000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/282450/bhx/snack-muc-vi-truyen-thong-talaethong-goi-100g_202507281356290014.jpg',18,1),
('Snack rong biển truyền thống Kimmy 6g', 15400, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/248296/bhx/thumb-248296_202411261128147817.jpg',4,1),
('Snack rong biển phô mai Kimmy 6g', 15400, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/248295/bhx/thumb-248295_202411261128107297.jpg',16,1),
('Snack mì ớt cấp độ 1 Enaak 28g', 6200, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/226028/bhx/thumb-226028_202411261127124889.jpg',18,1),
('Snack vị bò nướng Texas Lays 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/226184/bhx/snack-khoai-tay-vi-than-bo-nuong-texas-lays-wavy-goi-56g_202508211052036366.jpg',8,2),
('Snack vị tự nhiên Lays 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/193606/bhx/snack-poca-khoai-tay-tu-nhien-goi-52g_202510150939206616.jpg',5,2),
('Snack vị sườn BBQ Lays 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/193603/bhx/snack-poca-khoai-tay-suon-nuong-52g_202509241622313039.jpg',6,2),
('Snack vị tảo biển Lays 50g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/193602/bhx/thumb-193602_202411261120181132.jpg',8,2),
('Snack vị tảo biển OStar 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/79812/bhx/snack-ostar-rong-bien-48g_202511181412004013.jpg',10,2),
('Snack vị kim chi OStar 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/79804/bhx/snack-khoai-tay-ostar-vi-kim-chi-han-quoc-48g_202511181404222381.jpg',3,2),
('Snack vị bít tết Swing 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/79773/bhx/snack-swing-nysteak-48g_202511251647001007.jpg',5,2),
('Snack vị bò lúc lắc Poca 65g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/283009/bhx/thumb-283009_202411261130399401.jpg',6,2),
('Snack vị tảo biển Orion 56g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/305485/bhx/thumb-305485_202411261132325426.jpg',13,2),
('Snack vị bò Wagyu Lays Max 42g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/313108/bhx/thumb-313108_202411261134209919.jpg',8,2),
('Snack sốt Mayo kiểu Hàn OStar gói 50g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/327381/bhx/thumb-327381-1_202411261351494616.jpg',2,2),
('Snack tôm nướng ngũ vị Lays 42g', 22000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/341877/bhx/snack-khoai-tay-tom-hum-nuong-ngu-vi-lays-stax-lon-42g_202508041450197713.jpg',13,2),
('Snack vị tự nhiên Lays Stax 42g', 22000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/341873/bhx/snack-khoai-tay-vi-tu-nhien-lays-stax-lon-42g_202508041439104475.jpg',15,2),
('Snack vị kem chua & hành Lays 42g', 22000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/341658/bhx/snack-khoai-tay-vi-kem-chua-hanh-lays-stax-lon-42g_202508041427115966.jpg',1,2),
('Snack vị phô mai Cheddar Lays 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/339156/bhx/snack-khoai-tay-vi-pho-mai-cheddar-lays-wavy-goi-54g_202510150943599559.jpg',17,2),
('Snack vị tôm càng đút phô mai Lays 53g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/339146/bhx/snack-khoai-tay-vi-tom-cang-dut-lo-pho-mai-tan-chay-lays-goi-54g_202510150952082071.jpg',14,2),
('Snack mũ pháp sư vị phô mai Poca 58g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/339143/bhx/snack-hinh-mu-phap-su-vi-pho-mai-poca-goi-58g_202506041018517919.jpg',18,2),
('Snack mực lăn muối ớt Poca 60g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/339136/bhx/snack-muc-lan-muoi-ot-poca-goi-60g_202506031456283223.jpg',2,2),
('Snack hải sản vị cay ngọt Bento 22g', 29000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/338657/bhx/snack-hai-san-cay-ngot-bento-goi-22g_202505270844320036.jpg',6,2),
('Snack hải sản vị cay nồng Bento 22g', 29000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/338656/bhx/snack-hai-san-vi-cay-nong-bento-goi-22g_202505270857071401.jpg',4,2),
('Snack bắp vị phô mai cay Koikeya Karamucho 60g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/338644/bhx/snack-bap-vi-pho-mai-cay-koikeya-karamucho-goi-60g_202505261441423139.jpg',15,2),
('Snack bắp Koikeya Polinky 54g', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/338642/bhx/snack-bap-koikeya-polinky-goi-54g_202505261422442891.jpg',3,2),
('Bánh snack vị phô mai Nabati 80g', 24500, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/338613/bhx/banh-snack-vi-pho-mai-nabati-siip-hop-80g-4g-x-20-goi_202505260954160507.jpg',13,2),
('Bánh snack vị bắp rang Nabati 80g', 24500, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/338611/bhx/banh-snack-vi-bap-rang-nabati-siip-hop-80g-4g-x-20-goi_202505260943443524.jpg',7,2),
('Snack bắp vị socola Mplus 260g', 36000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/334868/bhx/snack-bap-vi-socola-mplus-chocolate-popcorn-chai-260g_202503281625537261.jpg',17,2),
('Snack hải sản Namprik Thái Bento 22g', 29000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/333971/bhx/thumb-333577-1_202501151003246540.jpg',10,2),
('Snack tôm xốt tương mật ong Oishi 70g', 12300, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/333564/bhx/thumb-333564-1_202412251002083319.jpg',10,2),
('Snack đậu xanh nước dừa Oishi 65g', 12300, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/333559/bhx/snack-dau-xanh-nuoc-dua-oishi-goi-70g_202506131528394847.jpg',15,2),
('Snack mực vị tương ớt Sriracha Talaethong 100g', 32000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/332876/bhx/snack-muc-vi-tuong-ot-sriracha-talaethong-goi-100g_202507281349450007.jpg',13,2),
('Snack bắp vị kem Mplus 260g', 36000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/327619/bhx/snack-bap-vi-kem-mplus-creamy-popcorn-chai-260g_202503281629438389.jpg',12,2),
('Snack bắp vị caramen Mplus 260g', 36000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/327616/bhx/thumb-327616_202411261403075655.jpg',8,2),
('Snack tôm vị cay nồng Oishi', 12000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/305937/bhx/snack-tom-vi-cay-nong-oishi-goi-70g_202505270833459443.jpg',3,2),
('Snack socola vị hazelnut Oishi Pillows 85g', 12300, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/305932/bhx/thumb-305932_202411261132580586.jpg',11,2),
('Snack kem Tiramisu Akiko Oishi 140g', 24000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/305735/bhx/thumb-305735_202411261132415506.jpg',8,2),
('Snack kem Mixed Berries Akiko Oishi 140g', 24000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/305733/bhx/thumb-305733_202411261132362500.jpg',19,2),
('Snack mực rong biển Talaethong 100g', 31000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/282757/bhx/snack-muc-vi-rong-bien-talaethong-goi-100g_202507281356252761.jpg',7,2),
('Snack mực Talaethong 100g', 31000, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/282450/bhx/snack-muc-vi-truyen-thong-talaethong-goi-100g_202507281356290014.jpg',18,2),
('Snack rong biển truyền thống Kimmy 6g', 15400, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/248296/bhx/thumb-248296_202411261128147817.jpg',4,2),
('Snack rong biển phô mai Kimmy 6g', 15400, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/248295/bhx/thumb-248295_202411261128107297.jpg',16,2),
('Snack mì ớt cấp độ 1 Enaak 28g', 6200, 'https://cdnv2.tgdd.vn/bhx-static/bhx/Products/Images/3364/226028/bhx/thumb-226028_202411261127124889.jpg',18,2);

-- ==========================================
-- 4. Order
-- ==========================================
CREATE TABLE order_table(
    OrderID UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    ShippingAddress TEXT NOT NULL,
    PaymentMethod VARCHAR(100) NOT NULL CHECK (PaymentMethod in ('MOMO','VNPAY')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    Status VARCHAR(50) NOT NULL CHECK (Status in ('ORDERED','PAYMENT_PENDING','PAID','FAILED')),
    BuyerID INT NOT NULL,
    FOREIGN KEY (BuyerID) REFERENCES users_table(UserID)
);
INSERT INTO order_table(ShippingAddress,PaymentMethod,Status,BuyerID)
VALUES
('130 Yen lang','VNPAY','PAID',1),
('36 Di An','MOMO','ORDERED',1),
('36 Di An','MOMO','ORDERED',1);
-- ==========================================
-- 5. OrderItem (Composite PK)
-- ==========================================
CREATE TABLE OrderItem (
    OrderID UUID NOT NULL,
    ProductID INT NOT NULL,
    Quantity INT NOT NULL,
    UnitPrice NUMERIC(12,2),
    PRIMARY KEY (OrderID, ProductID),
    FOREIGN KEY (OrderID) REFERENCES order_table(OrderID),
    FOREIGN KEY (ProductID) REFERENCES Product(ProductID)
);

-- ==========================================
-- 6. Review
-- ==========================================
CREATE TABLE Review (
    ReviewID SERIAL PRIMARY KEY,
    Content TEXT,
    Rating INT CHECK (Rating BETWEEN 1 AND 5),
    DatePosted TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    OrderID UUID NOT NULL,
    FOREIGN KEY (OrderID) REFERENCES order_table(OrderID)
);

-- ==========================================
-- 7. Promotion
-- ==========================================
CREATE TABLE Promotion (
    PromotionID SERIAL PRIMARY KEY,
    Name VARCHAR(150) NOT NULL,
    Type VARCHAR(10) NOT NULL CHECK (Type IN ('Shop', 'Product')),
    Value NUMERIC(12,2) NOT NULL,
    StartAt TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
    EndAt TIMESTAMP  DEFAULT CURRENT_TIMESTAMP,
    ShopID INT,
    ProductID INT,
    FOREIGN KEY (ShopID) REFERENCES Shop(ShopID),
    FOREIGN KEY (ProductID) REFERENCES Product(ProductID)
);

INSERT INTO Promotion(Name,Type,Value,ShopID)
VALUES
('Khuyến mãi cả shop','Product',0.2,1);
INSERT INTO Promotion(Name,Type,Value,ProductID)
VALUES
('Khuyến mãi sản phẩm Snack vị sườn BBQ Lays 53g','Product',0.3,3);


-- ==========================================
-- 8. SystemReport
-- ==========================================
CREATE TABLE SystemReport (
    ReportID SERIAL PRIMARY KEY,
    Title VARCHAR(200) NOT NULL,
    Description TEXT,
    Status VARCHAR(50) NOT NULL,
    ReportedByUserID INT NOT NULL,
    FOREIGN KEY (ReportedByUserID) REFERENCES users_table(UserID)
);
INSERT INTO SystemReport(Title,Description,Status,ReportedByUserID)
VALUES
('Báo cáo lỗi hệ thống','Hệ thống như ** ****','không vui',2);
-- ==========================================
-- 9. Wishlist
-- ==========================================
CREATE TABLE Wishlist (
    WishlistID SERIAL PRIMARY KEY,
    BuyerID INT UNIQUE NOT NULL,
    FOREIGN KEY (BuyerID) REFERENCES users_table(UserID)
);

-- ==========================================
-- 10. WishlistItem (Composite PK)
-- ==========================================
CREATE TABLE WishlistItem (
    WishlistID INT NOT NULL,
    ProductID INT NOT NULL,
    PRIMARY KEY (WishlistID, ProductID),
    FOREIGN KEY (WishlistID) REFERENCES Wishlist(WishlistID),
    FOREIGN KEY (ProductID) REFERENCES Product(ProductID)
);