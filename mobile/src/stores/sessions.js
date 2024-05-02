import { create } from "zustand"

const useSessionStore = create(set => ({
  isLogged: false,
  signIn: () => set({ isLogged: true }),
  signOut: () => set({ isLogged: false })
}))

export default useSessionStore
