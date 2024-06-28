import { productName, orderOfPriority, priceWithoutTaxes, priceWithTaxes, materials, categoryID, description, thumbnailURL, productImages } from "@@/lib/constants/labels"
import { string, number, array } from "yup"

export const nameValidator = string().min(2).max(255).label(productName)
export const orderOfPriorityValidator = number().positive().integer().label(orderOfPriority)
export const priceWithoutTaxesValidator = number().positive().label(priceWithoutTaxes)
export const priceWithTaxesValidator = number().positive().label(priceWithTaxes)
export const materialsValidator = string().min(2).max(255).label(materials)
export const categoryIDValidator = number().positive().integer().label(categoryID)
export const descriptionValidator = string().min(2).max(1500).label(description)
export const thumbnailURLValidator = string().min(2).max(255).label(thumbnailURL)