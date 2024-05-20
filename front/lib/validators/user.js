import { email, firstName, lastName, password } from "@@/lib/constants/labels"
import { string } from "yup"

export const firstNameValidator = string().min(2).max(255).label(firstName)
export const lastNameValidator = string().min(2).max(255).label(lastName)
export const emailValidator = string().email().label(email)
export const passwordValidator = string()
  .min(8)
  .matches(/(?=.*\p{Lu})(?=.*\p{Ll})(?=.*\d)(?=.*[^\d\p{L}]).*/u, "Must contain: 1 lower & 1 upper letters, 1 digit and 1 spe.char.")
  .label(password)
