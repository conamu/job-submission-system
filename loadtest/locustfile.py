from locust import HttpUser, task, between

class JobUser(HttpUser):
    wait_time = between(1, 2)

    HttpUser.job_id = None

    @task(1)
    def create_job(self):
        with self.client.post("/job", json={"payload": "hello from locust!"}, catch_response=True) as response:
            if response.status_code == 202:
                self.job_id = response.json().get("id")
            if response.status_code == 429:
                response.success()

    @task(5)
    def get_status(self):
        with self.client.get(f"/status/{self.job_id}", catch_response=True) as response:
            if response.status_code == 425:
                response.success()
            if response.status_code == 404:
                response.success()