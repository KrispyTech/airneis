import i18n from "i18next"
import { initReactI18next } from "react-i18next"
import * as ressources from "./ressources"

i18n.use(initReactI18next).init({
  compatibilityJSON: "v3",
  resources: Object.entries(ressources).reduce((acc, [key, value]) => ({ ...acc, [key]: { translation: value } }), {}),
  lng: "fr"
})

export default i18n
