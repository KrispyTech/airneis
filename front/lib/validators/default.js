import { number, string } from "yup"

export const stringValidator = string()
export const descriptionValidator = string().min(20)
export const idValidator = number()
export const numberValidator = number()
export const pageValidator = number().min(1).default(1)
