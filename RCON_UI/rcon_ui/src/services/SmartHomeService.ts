import store from "@/store";
import axios from "axios";

//TODO: turn this into system variable
//https://andrewp.online/smart/api/
//http://localhost:6969/smart/api/
const baseUri = "https://andrewp.online/smart/api/"

class SmartHomeService {
    async OpenGarage() {
      await axios.post(`${baseUri}open_garage`, null, { headers: authHeader()})
      .catch(e => {
        store.dispatch("addError", e.message)
      });
    }

    async GetSensorValues() {
      return await axios.get(`${baseUri}getSensorValues`, { headers: authHeader()})
      .then(res => {
        return res.data.Response;
      }) 
      .catch(e => {
        store.dispatch("addError", e.message);
      });

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

export default new SmartHomeService();