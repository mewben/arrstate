import { dispatch } from "@Store"
import { ACTIONS } from "@Enums"
import { requestApi } from "@Utils"

export const signIn = payload => {
  return dispatch({
    type: ACTIONS.AUTH.SIGNIN,
    payload: requestApi("/auth/signin2", "POST", {
      data: payload,
      noToken: true,
    }),
    meta: {
      globalError: false,
    },
  })
}

export const signIn2 = payload => {
  return {
    type: ACTIONS.AUTH.SIGNIN,
    payload2: fetch("http://localhost:5000/auth/signin", {
      method: "POST",
    }).then(response => response.json()),
    payload: requestApi("/auth/signin", "POST", {
      data: payload,
      noToken: true,
    }),
    meta: {
      globalError: false,
    },
  }
}

export const signIn23 = payload => (getState, dispatch) => {
  return dispatch({
    type: ACTIONS.AUTH.SIGNIN,
    payload: requestApi("/auth/signin", "POST", {
      data: payload,
      noToken: true,
    }),
    meta: {
      globalError: false,
    },
  })
}

export const signUp = payload => {
  return dispatch({
    type: ACTIONS.AUTH.SIGNUP,
    payload: requestApi("/auth/signup", "POST", {
      data: payload,
      noToken: true,
    }),
    meta: {
      globalError: false,
    },
  })
}
