import { MessageRequest } from "@/models/MessageRequest";
import store from "@/store";
import axios from "axios";

//TODO: turn this into system variable
//https://andrewp.online/RCON/api/admin/
//http://localhost:6969/RCON/api/admin/
const baseUri = "https://andrewp.online/RCON/api/admin/"

//TODO: replease with display on screen

class AdminCalloutService {
    GetPlayers() {
      return axios.get(`${baseUri}players`, { headers: authHeader()})
      .then(res => {
        return res.data.Response;
      })
      .catch(err => {
        store.dispatch("addError", err.message)
        return [];
      });
    }

    BroadCastMessage(message: string) {
      const request = new MessageRequest(message);
      axios.post(`${baseUri}broadcast`, request, { headers: authHeader()})
      .then(res => {
        return res;
      })
      .catch(err => {
        store.dispatch("addError", err.message)
        return err;
      });
    }

    async SaveAndRestart() {
      await axios.post(`${baseUri}save`, null, { headers: authHeader()})
      .catch(e => {
        store.dispatch("addError", e.message)
      });
    }

    async GetStatus() {
      const prom = await axios.get(`${baseUri}status`, { headers: authHeader()})
      .then(res => {
        return res.data.Response;
      })
      .catch(e => {
        return e.message
      });

      return prom
    }
}

function authHeader() {
    const token = localStorage.getItem("token");

    if (token) {
      return { Authorization: `Bearer ${token}` };
    } else {
      return {};
    }
  }

export default new AdminCalloutService();