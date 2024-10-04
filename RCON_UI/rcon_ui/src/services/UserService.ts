import { LoginRequest } from "@/models/LoginRequest";
import axios from "axios";

//TODO: turn this into system variable
//https://andrewp.online/RCON/api/
//http://localhost:6969/RCON/api/
const baseUri = "https://andrewp.online/RCON/api/"

class UserService {
    Login(request: LoginRequest) {
        return axios.post(`${baseUri}login`, request).then(res => {
            if (res.data.Token) {
                localStorage.setItem("token", res.data.Token);
            }
            return;
        })
    }

    Logout() {
        localStorage.removeItem("token")
    }
}

export default new UserService();