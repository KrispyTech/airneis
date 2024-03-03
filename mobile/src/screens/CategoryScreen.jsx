import { Image, ScrollView, Text, View } from "react-native"
import Page from "../components/Page"
import React, { useEffect, useState } from "react"
import { getCategoryById } from "../utils/categoriesServices"
import ErrorScreen from "./ErrorScreen"
import errors from "../utils/constants/errors"
import mobileConfig from "../../mobileConfig"
import ProductCard from "../components/ProductCard"

const CategoryScreen = props => {
  const {
    navigation,
    route: {
      params: { id }
    }
  } = props
  const [category, setCategory] = useState({
    products: []
  })

  const [hadError, setHadError] = useState(false)

  useEffect(() => {
    ;(async () => {
      try {
        const result = await getCategoryById(id)
        setCategory(result.data)
      } catch (e) {
        setHadError(true)
      }
    })()
  }, [id])

  if (hadError) {
    return <ErrorScreen navigation={navigation} text={errors.categories.unableToLoad} />
  }

  return (
    <Page navigation={navigation}>
      <ScrollView>
        <Image
          style={mobileConfig.images.fullScreenImageXl}
          source={{
            uri: category.thumbnailUrl
          }}
        ></Image>
        <View className="px-3">
          <Text className="text-4xl text-center font-bold py-5">{category.name}</Text>
          <Text className="text-lg font-medium uppercase">{category.description}</Text>
          <View className="py-8 ">
            {category.products.map(product => (
              <ProductCard product={product} key={product.id} />
            ))}
          </View>
        </View>
      </ScrollView>
    </Page>
  )
}

export default CategoryScreen
