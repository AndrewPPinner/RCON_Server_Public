import store from "@/store";
import axios from "axios";

//TODO: turn this into system variable
//https://andrewp.online/smart/api/
//http://localhost:6969/smart/api/
const baseUri = "https://andrewp.online/smart/api/"

class SmartHomeService {
    async OpenGarage() {
      await axios.post(`${baseUri}open_garage`, null, { headers: authHeader()})
      .catch((e: { message: any; }) => {
        store.dispatch("addError", e.message)
      });
    }

    async GetSensorValues() {
      return await axios.get(`${baseUri}getSensorValues`, { headers: authHeader()})
      .then((res: { data: { Response: any; }; }) => {
        return res.data.Response;
      }) 
      .catch((e: { message: any; }) => {
        store.dispatch("addError", e.message);
      });

    }

    async GetSensorDataGraph(location : string, type : string) : Promise<Array<any>> {
      const request =  {Location: location, Type: type};

      return await axios.post(`${baseUri}getSensorGraphData`, request, { headers: authHeader()})
      .then((res: { data: { Response: any; }; }) => {
        return res.data.Response;
      }) 
      .catch((e: { message: any; }) => {
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