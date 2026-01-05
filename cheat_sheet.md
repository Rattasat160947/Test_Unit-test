# GoValidator Cheat Sheet (สำหรับสอบ)

--------------------------------------------------------------------------------
[ 1. พื้นฐาน & การตรวจสอบค่า (Common) ] ⭐ ใช้บ่อยที่สุด
--------------------------------------------------------------------------------
Tag: required           -> ห้ามเป็นค่าว่าง / nil / 0
Tag: isnull             -> ต้องเป็นค่าว่าง (Is_null)
Tag: notnull            -> ต้องไม่เป็นค่าว่าง (IsNotNull)
Tag: bool               -> ต้องเป็น string ที่แปลงเป็น boolean ได้ ("true", "false")

--------------------------------------------------------------------------------
[ 2. ตัวเลข & การคำนวณ (Numeric) ]
--------------------------------------------------------------------------------
Tag: numeric            -> เป็นตัวเลขล้วนๆ (String) "123456"
Tag: int                -> เป็นจำนวนเต็ม (Integer)
Tag: float              -> เป็นทศนิยม (Float)
Tag: range(min|max)     -> ค่าอยู่ระหว่าง min-max
Tag: natural            -> จำนวนนับ (Positive Integer)
Tag: negative           -> ค่าติดลบ (< 0)
Tag: positive           -> ค่าบวก (> 0)
Tag: nonpositive        -> ค่า <= 0
Tag: nonnegative        -> ค่า >= 0
Tag: whole              -> จำนวนเต็ม (ไม่มีทศนิยม)
Tag: divisibleby(n)     -> หารด้วย n ลงตัว

--------------------------------------------------------------------------------
[ 3. ตัวอักษร & ข้อความ (String & Text) ]
--------------------------------------------------------------------------------
Tag: alpha              -> ตัวอักษรล้วน (a-z, A-Z)
Tag: alphanum           -> ตัวอักษร + ตัวเลข
Tag: ascii              -> อักขระ ASCII
Tag: printableascii     -> ASCII ที่พิมพ์ได้
Tag: multibyte          -> ตัวอักษรหลาย Byte (เช่น ภาษาไทย/จีน/ญี่ปุ่น)
Tag: lowercase          -> ตัวพิมพ์เล็กทั้งหมด
Tag: uppercase          -> ตัวพิมพ์ใหญ่ทั้งหมด
Tag: fullwidth          -> ตัวอักษรแบบ Full-width
Tag: halfwidth          -> ตัวอักษรแบบ Half-width
Tag: variablewidth      -> ผสมทั้ง Full และ Half

Tag: stringlength(n|m)  -> ความยาวตัวอักษร n ถึง m
Tag: runelength(n|m)    -> ความยาวตัวอักษร (นับ Unicode/ภาษาไทย เป็น 1 ตัว) ⭐ แนะนำอันนี้ถ้าใช้ภาษาไทย
Tag: in(a|b|c)          -> ต้องมีค่าตรงกับ a หรือ b หรือ c
Tag: matches(regex)     -> ตรงตาม Regular Expression

--------------------------------------------------------------------------------
[ 4. อินเทอร์เน็ต & เครือข่าย (Network) ]
--------------------------------------------------------------------------------
Tag: email              -> รูปแบบอีเมล
Tag: url                -> รูปแบบ URL (http/https)
Tag: requrl             -> Request URL
Tag: requri             -> Request URI
Tag: datauri            -> Data URI scheme
Tag: dialstring         -> เบอร์โทรศัพท์/Dial String

Tag: ip                 -> IP Address (v4 หรือ v6)
Tag: ipv4               -> IP Address v4
Tag: ipv6               -> IP Address v6
Tag: cidr               -> CIDR notation (e.g. 192.168.0.0/24)
Tag: mac                -> MAC Address
Tag: port               -> Port number
Tag: dns                -> DNS Name / Hostname
Tag: host               -> Hostname (DNS หรือ IP)

--------------------------------------------------------------------------------
[ 5. รหัส & รูปแบบเฉพาะ (Code & Formats) ]
--------------------------------------------------------------------------------
Tag: uuid               -> UUID (Version ไหนก็ได้)
Tag: uuidv3             -> UUID Version 3
Tag: uuidv4             -> UUID Version 4
Tag: uuidv5             -> UUID Version 5
Tag: ulid               -> ULID
Tag: mongo              -> MongoDB ObjectID

Tag: creditcard         -> เลขบัตรเครดิต
Tag: isbn10             -> เลขหนังสือ ISBN-10
Tag: isbn13             -> เลขหนังสือ ISBN-13
Tag: ssn                -> Social Security Number (US)
Tag: semver             -> Semantic Version (x.y.z)
Tag: json               -> ข้อความรูปแบบ JSON
Tag: base64             -> ข้อความ Base64

--------------------------------------------------------------------------------
[ 6. สี & การเข้ารหัส (Color & Hashing) ]
--------------------------------------------------------------------------------
Tag: hexcolor           -> รหัสสี Hex (#RRGGBB)
Tag: rgbcolor           -> รหัสสี RGB (rgb(r,g,b))
Tag: hexadecimal        -> เลขฐาน 16
Tag: md4                -> MD4 Hash
Tag: md5                -> MD5 Hash
Tag: sha1               -> SHA1 Hash
Tag: sha256             -> SHA256 Hash
Tag: sha384             -> SHA384 Hash
Tag: sha512             -> SHA512 Hash
Tag: ripemd128          -> RipeMD128
Tag: ripemd160          -> RipeMD160
Tag: tiger128           -> Tiger128
Tag: tiger160           -> Tiger160
Tag: tiger192           -> Tiger192

--------------------------------------------------------------------------------
[ 7. วันที่ & เวลา (Date & Time) ]
--------------------------------------------------------------------------------
Tag: rfc3339            -> วันที่แบบ RFC3339
Tag: rfc3339withoutzone -> RFC3339 แบบไม่มี Timezone
Tag: unix               -> Unix Timestamp
Tag: time(format)       -> ตรวจสอบวันที่ตาม Format ที่กำหนด
                           เช่น time(02-Jan-2006)

--------------------------------------------------------------------------------
[ 8. พิกัด & สถานที่ (Location) ]
--------------------------------------------------------------------------------
Tag: latitude           -> ละติจูด
Tag: longitude          -> ลองจิจูด
Tag: iso3166alpha2      -> รหัสประเทศ 2 ตัว (US, TH)
Tag: iso3166alpha3      -> รหัสประเทศ 3 ตัว (USA, THA)

*/