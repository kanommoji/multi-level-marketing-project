# multi-level-marketing-project

## ข้อตกลง Name Convention
1. การตั้งชื่อเทส ให้ตั้งในรูปแบบ Snake Case โดยประกอบด้วยชื่อฟังก์ชันที่ทดสอบ Input และ Output
เชื่อมชื่อฟังก์ชันกับ Input ด้วย By และเชื่อมกับ Output ด้วย Should_Be <br>
ยกตัวอย่างเช่น Test_CheckCondition_By_Member_Level_6_MonthlyPoint_400_TeamPoint_20050_TeamMemberMoreEmerald_2_Should_Be_True()


## ข้อตกลงใน Commit message
ADD : เพิ่ม สร้าง <br>
DELETE : ลบ<br>
EDIT : แก้ไข ปรับปรุง<br>

Example : `"ADD : สร้างไฟล์ README.md และเพิ่มข้อตกลงในการทำงาน"`

## ข้อตกลงในการเขียน Error Test Message

    if expectedResult != actualResult {
            t.Errorf("Expect %v but get %v", expectedResult, actualResult)
        }
    
## ข้อตกลงในการจัดรูปแบบ Function Unit Test
ให้ทำการ เว้นบรรทัด แยกการทำงานออกเป็น 3 ส่วนคือ 1.arrange 2.action 3.assert ตัวอย่าง

    func Test_CheckCondition_By_Member_Level_6_MonthlyPoint_400_TeamPoint_20050_TeamMemberMoreEmerald_2_Should_Be_True(t *testing.T) {
        expectedResult := true
        detailsMemberAlpha := MemberAlpha{
            Level: 6,
            TeamPoint: 20050,
            MonthlyPoint: 400,
            TeamMember: TeamMember{
                HigherEmerald: 2,
            },
        }

        actualResult := checkCondition(detailsMemberAlpha)
        
        if expectedResult != actualResult {
            t.Errorf("Expect %v but get %v", expectedResult, actualResult)
        }
    }
    
## Project Setup
เริ่มต้นสร้าง user ทำหรับ database โดยเข้า mysql ด้วยสิทธิ์ root และมอบสิทธิ์ให้กับ user ที่สร้างโดยคำสั่ง

    CREATE USER 'mlm_dev'@'localhost' IDENTIFIED BY 'mlm_dev';
    GRANT ALL PRIVILEGES ON *.* TO 'mlm_dev'@'localhost' IDENTIFIED BY 'mlm_dev';
    
ค่า Setup Config เกี่ยวกับ Connection ของ database สามารถแก้ไขได้ที่ไฟล์ config/development.yml

หรือสามารถใช้ค่า Environment แทนได้ โดยกำหนดค่าก่อนจะเริ่มรัน main ยกตัวอย่าง

    $ export USERNAME_DB=root PASSWORD_DB=root DATABASE_NAME=mlm HOST_DB=127.0.0.1 PORT_DB=3306
    $ go build cmd/main/main.go
    $ ./main.exe



## Unit Test ,  Integration Test

    $ mysql -umlm_dev -pmlm_dev < sql/prepared-data.sql
    $ TIME=20190701 go test ./...
    $ mysql -umlm_dev -pmlm_dev < sql/drop-mlm-database.sql
    
## ATTD
    
    $ TIME=20190701 go run cmd/main/main.go
    $ mysql -umlm_dev -pmlm_dev < sql/prepared-data.sql
    $ newman run atdd/api/promote-member-level.json
    $ mysql -umlm_dev -pmlm_dev < sql/drop-mlm-database.sql


## วิธีการใช้งาน API
เริ่มต้น Start main

    $ go build cmd/main/main.go
    $ ./main.exe
    
หรือ

    $ go run cmd/main/main.go
    
    

URL : https://localhost:8080/action_point
Method : POST

ตัวอย่าง Request Body

    {
        "user_referral":10029,
        "new_point":50
    }

ตัวอย่าง Response Body

    {
        "member_id": 10029,
        "member_name": "ชนา",
        "leader_id": 20029,
        "level": 7,
        "my_point": 1050,
        "monthly_point": 20050,
        "team_point": 400
    }

**Request Body**

| Key | Description |
|--|--|
| user_referral | ID ของสมาชิกที่หาลูกค้า |
| new_point | Point ที่ได้รับ |

**Response Body** 

| Key | Description |
|--|--|
| member_id | ID ของสมาชิก |
| member_name | ชื่อสมาชิก |
| leader_id | ID ของหัวหน้าทีม |
| level | ตำแหน่ง |
| my_point | คะแนนของสมาชิก |
| monthly_point | คะแนนสะสมผลงานรายเดือน |
| team_point | คะแนนของทีม |
