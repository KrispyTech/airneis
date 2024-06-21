import { productName, priority, priceET, priceIT, materials, category, description } from "@@/lib/constants/labels"
import { string, number } from "yup"

export const nameValidator = string().min(2).max(255).label(productName)
export const priorityValidator = number().positive().integer().label(priority)
export const priceETValidator = number().positive().label(priceET)
export const priceITValidator = number().positive().label(priceIT)
export const materialsValidator = string().min(2).max(255).label(materials)
export const categoryValidator = string().min(2).max(255).label(category)
export const descriptionValidator = string().min(2).max(1500).label(description)