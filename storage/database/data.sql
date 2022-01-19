USE contact;

INSERT INTO hosts (name, url, email, subject) VALUES ('TEST', 'test.com', 'automailr.noreply@gmail.com', 'Test Contact Form');
INSERT INTO fields (name, required, host_id) VALUES ('required', 1, 1), ('not_required', 0, 1);
/* Must copy template from ./email-templates */
INSERT INTO templates (template, host_id) VALUES ('<!DOCTYPE html>
<html>
  <head></head>
  <body>
    <div>
      <ul>
        <li>Required: {{.required}}</li>
        {{ if .not_required }}
        <li>Not Required: {{.not_required}}</li>
        {{end}}
      </ul>
    </div>
  </body>
</html>', 1);

INSERT INTO hosts (name, url, email, subject) VALUES ('Oakville Windows & Doors', 'owd.noahvarghese.me', 'info@oakvillewindows.com', 'Web Request');
INSERT INTO fields (name, required, host_id) VALUES ('name', 1, 2), ('email', 1, 2), ('phone', 0, 2), ('city', 0, 2), ('message', 1, 2), ('products', 0, 2);
/* Must copy template from ./email-templates */
INSERT INTO templates (template, host_id) VALUES ('<!DOCTYPE html>
<html>
  <head></head>
  <body>
    <div>
      <div>
        <sub
          ><em
            >Please do not reply to this email. It will not reach the intended
            recipient. If there are any issues please email
            <a href="mailto:varghese.noah@gmail.com">Noah Varghese</a></em
          ></sub
        >
      </div>
      <ul>
        <li>Name: {{.name}}</li>
        <li>Email: {{.email}}</li>
        {{ if .Phone }}
        <li>Phone: {{.Phone}}</li>
        {{end}} {{ if .City }}
        <li>City: {{.City}}</li>
        {{end}} {{ if .Products }}
        <li>City: {{.Products}}</li>
        {{end}}
      </ul>
      <p>{{.message}}</p>
    </div>
  </body>
</html>', 2);
