import { LoginRequest } from "@/models/LoginRequest";
import UserService from "@/services/UserService";
import { createStore } from "vuex";

export default createStore({
  state: {
    isAuthorized: false,
    token: "" as string | null,
    errors: new Map<number, ErrorObject>(),
    errorsThisSession: 0
  },
  getters: {},
  mutations: {
    loginSuccess(state, token) {
      state.isAuthorized = true;
      state.token = token;
    },
    loginFailure(state, error) {
      state.isAuthorized = false;
      state.token = null;
      if (error.response?.data.Response != null) {
        state.errors.set(state.errorsThisSession++, new ErrorObject(error.response.data.Response, new Date(new Date().getTime() + 6*1000)))
      } else {
        state.errors.set(state.errorsThisSession, new ErrorObject("Unable to connect with server. Please try again later.", new Date(new Date().getTime() + 4*60*1000)))
      }
    },
    logout(state) {
      state.isAuthorized = false;
      state.token = null;
    },
    addError(state, error) {
      state.errors.set(state.errorsThisSession++, new ErrorObject(error, new Date(new Date().getTime() + 6*1000)));
    },
    removeError(state, id) {
      state.errors.delete(id);
    }
  },
  actions: {
    login({ commit }, request: LoginRequest) {
      return UserService.Login(request).then(
        res => {
          commit("loginSuccess", res);
          return Promise.resolve(res);
        },
        error => {
          commit("loginFailure", error);
          return Promise.reject(error);
        }
      );
    },
    logout({ commit }) {
      UserService.Logout();
      commit("logout");
    },
    addError({ commit }, error: string) {
      commit("addError", error);
    },
    removeError({ commit }, id: number) {
      commit("removeError", id);
    }
  },
  modules: {},
});

class ErrorObject {
  Error: string
  Time: Date

  constructor(error: string, time: Date) {
    this.Error = error;
    this.Time = time
  }
}
