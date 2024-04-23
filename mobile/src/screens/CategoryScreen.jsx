import { Button, Image, ScrollView, Text, TouchableHighlight, View } from "react-native"
import Page from "../components/Page"
import React, { useEffect, useState } from "react"
import { getCategoryById } from "../utils/categoriesServices"
import ErrorScreen from "./ErrorScreen"
import errors from "../utils/constants/errors"
import mobileConfig from "../../mobileConfig"
import ProductCard from "../components/ProductCard"
import screens from "../utils/constants/screens"
import { useTranslation } from "react-i18next"

const CategoryScreen = props => {
  const { t } = useTranslation()
  const {
    navigation,
    route: {
      params: { id }
    }
  } = props

  const [category, setCategory] = useState({
    products: []
  })

  const [hasError, setError] = useState(false)

  useEffect(() => {
    ;(async () => {
      try {
        const { data: result } = await getCategoryById(id)
        setCategory(result)
      } catch (e) {
        setError(true)
      }
    })()
  }, [id])

  if (hasError) {
    return <ErrorScreen navigation={navigation} text={t(errors.category.unableToLoad)} />
  }

  const navigateToProduct = productId => () => navigation.navigate(screens.productPage, { productId })

  return (
    <Page navigation={navigation}>
      <ScrollView className="relative">
        <View>
          <Image
            style={mobileConfig.images.fullScreenImageXL}
            source={{
              url: category.thumbnailURL
            }}
          ></Image>
          <Text className="text-4xl text-center font-bold absolute top-1/2 left-0 right-0 text-white" {...mobileConfig.dropShadow.textOutline}>
            {category.name}
          </Text>
        </View>
        <View className="px-3 py-8">
          <Text className="text-lg font-medium uppercase">{category.description}</Text>
          <View className="py-8 gap-5">
            {category.products &&
              category.products.map(product => (
                <TouchableHighlight key={`product${product.ID}`} onPress={navigateToProduct(product.ID)}>
                  <View>
                    <ProductCard product={product} />
                  </View>
                </TouchableHighlight>
              ))}
          </View>
        </View>
      </ScrollView>
    </Page>
  )
}

export default CategoryScreen
