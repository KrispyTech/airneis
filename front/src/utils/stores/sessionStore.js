import { createSession, deleteSession } from "@/utils/services/sessionServices"
import { create } from "zustand"
import routes from "@@/lib/constants/routes"
import { jwtDecode } from "jwt-decode"

const initialValues = {
  firstName: null,
  lastName: null,
  email: null,
  isAdmin: false,
  isActive: false
}

const useSessionStore = create(set => ({
  ...initialValues,
  createUserSession: (inputEmail, inputPassword, router) =>
    set(async () => {
      const {
        data: { data: token }
      } = await createSession(inputEmail, inputPassword)

      const {
        payload: { userData: firstName, lastName, email, isAdmin, isActive }
      } = jwtDecode(token)

      router.push(routes.home)

      return {
        firstName,
        lastName,
        email,
        isAdmin,
        isActive
      }
    }),
  removeUserSession: () =>
    set(async () => {
      await deleteSession()

      return initialValues
    })
}))

export default useSessionStore
