Запуск проекта:
make docker-build
docker-compose up --build
API:
POST /employees
DELETE /employees/{employeeId}
GET /employees/company/{companyId}
GET /employees/{companyId}/{departmentName}
PUT /employees/{employeeId}
Добавление сотрудников
   - Метод: POST
   - Описание: Позволяет добавить нового сотрудника.  
   - Ответ: Возвращает Id добавленного сотрудника.
   - Пример запроса:
  curl -X POST -H "Content-Type: application/json" -d '{
  "Name": "Иван",
  "Surname": "Иванов",
  "Phone": "+79991112233",
  "CompanyId": 123,
  "Passport": {
    "Type": "Паспорт РФ",
    "Number": "1234567890"
  },
  "Department": {
    "Name": "Отдел продаж",
    "Phone": "+74951234567"
  }
}' http://localhost:8080/employees
Удаление сотрудника:
  - Метод: DELETE
   - Описание: Удаляет сотрудника по заданному Id.
   - Пример запроса:
  curl -X DELETE http://localhost:8080/employees/1
Вывод списка сотрудников для указанной компании:
    - Метод: GET
   - Описание: Возвращает список сотрудников для указанной компании. Все доступные поля.
   - Пример запроса:
  curl http://localhost:8080/employees/company/123
Вывод списка сотрудников по отделу компании
   - Метод: GET
   - Описание: Возвращает список сотрудников для указанного отдела компании. Все доступные поля.
   - Пример запроса:
  curl http://localhost:8080/employees/company/123/HR
Изменение данных сотрудника
   - Метод: PATCH
   - Описание: Позволяет изменить информацию о сотруднике по его Id. Изменения применяются только к указанным полям.
   - Пример запроса:
  curl -X PUT -H "Content-Type: application/json" -d '{
  "Phone": "+79999999999"
}' http://localhost:8080/employees/1
