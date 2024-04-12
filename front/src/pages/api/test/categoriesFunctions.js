/* eslint-disable no-console */
import { createCategory, deleteCategory, getCategories, getCategoryById, updateCategory } from "@/utils/services/CategoryServices"

const SAMPLE_CATEGORY = {
  ID: 127,
  name: "Test",
  thumbnailURL: "https://picsum.photos/id/48/200/300",
  slug: "test",
  orderOfDisplay: 127,
  description:
    "Imperdiet veniam aliquip gubergren illum incidunt nonumy nonummy clita doming consetetur. Gubergren sanctus reprehenderit culpa laborum iure amet iriure labore accusam invidunt. Veniam mazim fugiat. Amet vulputate magna."
}

// BEFORE STARTING THE TEST BE SURE THAT YOU DO NOT HAVE A CATEGORY THAT HAVE 127 AS ORDER OF DISPLAY OR 127 AS ID!!
const handler = async (req, res) => {
  console.log("Welcome to the categories fuction test")
  console.log("_____Testing POST_____")

  try {
    const category = await createCategory(SAMPLE_CATEGORY)
    console.log("Result of post category", category)
  } catch (e) {
    // As it is a test endpoint we need to have some debug logs
    console.log("Post category failed", e.message, e.path)
    res.status(500).send("Failed at POST")

    return
  }

  console.log("_____Testing GET_____")

  try {
    const categories = await getCategories()
    console.log("Result of get categories", categories)
  } catch (e) {
    console.log("Get categories failed", e)

    res.status(500).send("Failed at getCategory")

    return
  }

  try {
    const category = await getCategoryById(SAMPLE_CATEGORY.ID)
    console.log("Result of get category", category)
  } catch (e) {
    console.log("Get category failed", e)

    res.status(500).send("Failed at getCategory")

    return
  }

  console.log("_____Testing PATCH_____")

  try {
    const updateCategoryInput = {
      name: "Updated category"
    }

    const updateCategoryResult = await updateCategory(SAMPLE_CATEGORY.ID, updateCategoryInput)
    console.log("Result of get categories", updateCategoryResult)
  } catch (e) {
    console.log("Get categories failed", e)

    res.status(500).send("Failed at updateCategory")

    return
  }

  console.log("_____Testing DELETE_____")

  try {
    const deletedCategory = await deleteCategory(SAMPLE_CATEGORY.ID)
    console.log("Result of get categories", deletedCategory)
  } catch (e) {
    console.log("Get categories failed", e)

    res.status(500).send("Failed at deleteCategory")

    return
  }

  res.status(200).send("No errors")
}

export default handler
