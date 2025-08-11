Candidate Name: Chong Shao Yang

## Part 3: Database

a)

i) SQL query to count total employee for each department code

```sql
SELECT COUNT(e.ID), d.Code
FROM Employee e
INNER JOIN Department d
  ON d.ID = e.DepartmentID
GROUP BY d.Code
```

ii)

```sql
SELECT e.Name, e.Salary
FROM Employee e
INNER JOIN Department d
  ON d.ID = e.DepartmentID
WHERE e.Salary BETWEEN 3000 AND 4000
ORDER BY d.Code, e.Name
```

---

b)

i)

_bare minimum database design_

### Table: Subject

| Key | Column  | Type         | Description                  |
| --- | ------- | ------------ | ---------------------------- |
| PK  | ID      | INT          | Primary Key                  |
|     | Subject | VARCHAR(255) | (eg. Math, Science, English) |

### Table: Student

| Key | Column    | Type         | Description          |
| --- | --------- | ------------ | -------------------- |
| PK  | ID        | INT          | Primary key          |
|     | FirstName | VARCHAR(255) | Student's first name |
|     | LastName  | VARCHAR(255) | Student's last name  |
|     | StudentID | VARCHAR(255) | Student's unique id  |

### Table: Semester

| Key | Column | Type | Description                   |
| --- | ------ | ---- | ----------------------------- |
| PK  | ID     | INT  | Primary key                   |
|     | Year   | INT  | Unique year number (eg. 2025) |
|     | SemNum | INT  | Semester number (1 to 4)      |

### Table: Class

| Key | Column     | Type         | Description                  |
| --- | ---------- | ------------ | ---------------------------- |
| PK  | ID         | INT          | Primary Key                  |
| FK  | SemesterID | INT          | Reference to **Semester.ID** |
| FK  | SubjectID  | INT          | Reference to **Subject.ID**  |
|     | Name       | VARCHAR(255) | Class name                   |

### Table: Enrollment

| Key | Column       | Type     | Description                 |
| --- | ------------ | -------- | --------------------------- |
| PK  | ID           | INT      | Primary Key                 |
| FK  | StudentID    | INT      | Reference to **Student.ID** |
| FK  | ClassID      | INT      | Reference to **Class.ID**   |
|     | RegisteredAt | DATETIME | Registered time             |

```sql
CREATE TABLE Subject (
    ID INT PRIMARY KEY,
    Subject VARCHAR(255) NOT NULL
);

CREATE TABLE Student (
    ID INT PRIMARY KEY,
    FirstName VARCHAR(255) NOT NULL,
    LastName VARCHAR(255) NOT NULL,
    StudentID VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE Semester (
    ID INT PRIMARY KEY,
    "Year" INT NOT NULL,
    SemNum INT NOT NULL CHECK (SemNum BETWEEN 1 AND 4),
    CONSTRAINT uniq_year_code UNIQUE ("Year", SemNum)
);

CREATE TABLE Class (
    ID INT PRIMARY KEY,
    SemesterID INT NOT NULL,
    SubjectID INT NOT NULL,
    "Name" VARCHAR(255) NOT NULL,
    FOREIGN KEY (SemesterID) REFERENCES Semester(ID),
    FOREIGN KEY (SubjectID) REFERENCES Subject(ID),
    CONSTRAINT uniq_sem_subject_name UNIQUE (SemesterID, SubjectID, "Name")
);

CREATE TABLE Enrollment (
    ID INT PRIMARY KEY,
    StudentID INT NOT NULL,
    ClassID INT NOT NULL,
    RegisteredAt DATETIME NOT NULL,
    FOREIGN KEY (StudentID) REFERENCES Student(ID),
    FOREIGN KEY (ClassID) REFERENCES Class(ID)
);

-- Trigger to check semester count per year
DELIMITER $$

CREATE TRIGGER trig_check_sem_count
BEFORE INSERT ON Semester
FOR EACH ROW
BEGIN
  DECLARE sem_count INT;
  SELECT COUNT(*) INTO sem_count
  FROM Semester
  WHERE Year = NEW.Year;

  IF sem_count >= 4 THEN
    SIGNAL SQLSTATE '45000'
    SET MESSAGE_TEXT = 'Only 4 semesters are allowed for a year';
  END IF;
END$$

-- Trigger to check 2 max classes per subject per sem
CREATE trig_check_class_count
BEFORE INSERT ON Class
FOR EACH ROW
BEGIN
  DECLARE class_count INT;
  SELECT COUNT(*) INTO class_count
  FROM Class
  WHERE SemesterID = New.SemesterID
    AND SubjectID = New.SubjectID;

  IF class_count >= 2 THEN
    SIGNAL SQLSTATE '45000'
    SET MESSAGE_TEXT = 'Only 2 classes per subject are allowed per year';
  END IF;
END$$

DELIMITER;
```
