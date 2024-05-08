import { createSession, deleteSession } from "@/utils/services/sessionServices"
import { create } from "zustand"

const useSessionStore = create(set => ({
  createUserSession: (email, password) => set(() => createSession(email, password)),
  removeUserSession: () => set(deleteSession)
}))

export default useSessionStore
