import {Post} from "./apiBase";

export const login = (username, password) => {
  return Post("/auth", null, {username: username, password: password});
};

export const signOut = () => {

};