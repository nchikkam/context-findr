import axios from "axios";
import { Server } from "../common/helper";

class AuthService {
  login(email: string, password: string) {
    console.log("Trying to Login....")
    return axios
      .post(Server.baseURL + "/signin", {
        email,
        password
      })
      .then(response => {
        console.log(response)
        if (response.data.token) {
          localStorage.setItem("user", JSON.stringify(response.data));
        }

        return response.data;
      });
  }

  logout() {
    localStorage.removeItem("user");
  }

  register(username: string, email: string, password: string) {
    return axios.post(Server.baseURL + "/register", {
      username,
      email,
      password
    });
  }

  getCurrentUser() {
    const userStr = localStorage.getItem("user");
    if (userStr) return JSON.parse(userStr);

    return null;
  }
}

const service = new AuthService();

export default service;