import { useEffect, useState } from "react"
import Page from "../components/Page"
import errors from "../utils/constants/errors"
import ErrorScreen from "./ErrorScreen"
import { getProductById } from "../utils/services/productServices"
import { Text, TouchableHighlight, View } from "react-native"
import ImageGallery from "../components/ImageGallery"
import ProductCard from "../components/ProductCard"

// Placeholder to be removed when the images will be ready
const sampleImage = ["https://picsum.photos/400/300", "https://picsum.photos/400/400"]
const ProductScreen = props => {
  const {
    navigation,
    route: {
      params: { id }
    }
  } = props

  const [product, setProduct] = useState({})
  const [imageToDisplay, setImageToDisplay] = useState(0)
  const [hasError, setError] = useState(false)
  const handleChangeImage = i => () => {
    setImageToDisplay(i)
  }

  const handleAddToCart = () => {
    // To do
  }

  useEffect(() => {
    ;(async () => {
      try {
        const { data: result } = await getProductById(id)
        setProduct(result)
      } catch (e) {
        setError(true)
      }
    })()
  }, [id])

  if (hasError) {
    return <ErrorScreen navigation={navigation} text={errors.product.unableToLoad} />
  }

  return (
    <Page navigation={navigation}>
      <View>
        <ImageGallery images={sampleImage} imageToDisplay={imageToDisplay} handleChangeImage={handleChangeImage} />
        <View className="mx-3 pb-6">
          <View className="flex-row justify-between mt-3">
            <Text className="text-2xl font-extrabold uppercase">{product.name}</Text>
            <Text className="text-2xl font-extrabold">{product.priceWithTaxes}€</Text>
          </View>
          <Text className="text-darkGray text-xl font-bold">{product.outOfStock ? "Out Of Stock" : "In stock"}</Text>
          <View className="flex-row gap-4 pt-2">
            {/* Remove check when backoffice is merge and handle material*/}
            {product.materials &&
              product.materials.map((material, i) => (
                <View key={`material-${i}`} className="bg-lightGray rounded-full border px-2">
                  <Text className="text-lg">{material}</Text>
                </View>
              ))}
          </View>
          <Text className="text-lg pt-2 font-semibold">{product.description}</Text>
          <TouchableHighlight title="add to cart" className="border-4 bg-red-400 mx-auto px-6 py-3 mt-2" onPress={handleAddToCart}>
            <Text className="text-xl font-bold">AJOUTER AU PANIER</Text>
          </TouchableHighlight>
          <Text className="mx-auto text-3xl font-extrabold uppercase mt-4 mb-4">Produits Similaires</Text>
          {/* Remove check when api will handle similar products */}
          {product.similarProducts &&
            product.similarProducts.map(productSingle => <ProductCard key={`produitSimilaire-${productSingle.id}`} product={productSingle} />)}
        </View>
      </View>
    </Page>
  )
}

export default ProductScreen
