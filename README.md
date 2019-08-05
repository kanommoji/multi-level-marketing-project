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
    
## คำสั่งในการทดสอบ


## Unit Test ,  Intergration Test

    $  mysql -uroot -proot < sql/drop-mlm-database.sql
    $ TIME=20190701 go test ./...
    $ mysql -uroot -proot < sql/drop-mlm-database.sql
    
## ATTD
    
    $ TIME=20190701 go run cmd/main/main.go
    $ mysql -uroot -proot < sql/prepared-data.sql
    $ newman run atdd/api/promote-member-level.json
    $ mysql -uroot -proot < sql/drop-mlm-database.sql


## วิธีการใช้งาน API
เริ่มต้น Start main

    $ go build cmd/main/main.go
    $ ./main.exe -http=8080

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

## คำสั่งตั้งค่า Config 
    $ export USERNAME=root PASSWORD=root DATABASE=mlm HOST=127.0.0.1 PORT=3306
