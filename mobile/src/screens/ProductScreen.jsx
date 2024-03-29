import { useEffect, useState } from "react"
import Page from "../components/Page"
import errors from "../utils/constants/errors"
import ErrorScreen from "./ErrorScreen"
import { getProductById } from "../utils/productServices"
import { Text, TouchableHighlight, View } from "react-native"
import ImageGallery from "../components/ImageGallery"
import ProductCard from "../components/ProductCard"

const productEx = {
  id: 1,
  name: "Canapé Blanc",
  thumbnailURL: "https://picsum.photos/400/300",
  outOfStock: false,
  materials: ["Wood", "Coton"],
  produitsSimilaires: [{ name: "Canapé Marron", id: 2, thumbnailURL: "https://picsum.photos/400/300", priceWithTaxes: 1450 }],
  description:
    "Lorem ipsum dolor sit amet. Aut nihil Quis et dolorum quis hicsint voluptas id aspernatur corporis.ad voluptatem aliquam. At autem molestias et sequi harum aut rerum quia est  consequatur cupiditate et aliquid rerum et labore  dignissimos?",
  images: ["https://picsum.photos/400/300", "https://picsum.photos/400/400"],
  priceWithTaxes: 1200
}

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
        // const result = await getProductById(id)
        // setProduct(result.data)
        setProduct(productEx)
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
        <ImageGallery images={productEx.images} imageToDisplay={imageToDisplay} handleChangeImage={handleChangeImage} />
        <View className="mx-3 pb-6">
          <View className="flex-row justify-between mt-3">
            <Text className="text-2xl font-extrabold uppercase">{product.name}</Text>
            <Text className="text-2xl font-extrabold">{product.priceWithTaxes}€</Text>
          </View>
          <Text className="text-darkGray text-xl font-bold">{productEx.outOfStock ? "Out Of Stock" : "In stock"}</Text>
          <View className="flex-row gap-4 pt-2">
            {productEx.materials.map((material, i) => (
              <View key={`material-${i}`} className="bg-lightGray rounded-full border px-2">
                <Text className="text-lg">{material}</Text>
              </View>
            ))}
          </View>
          <Text className="text-lg pt-2 font-semibold">{productEx.description}</Text>
          <TouchableHighlight title="add to cart" className="border-4 bg-red-400 mx-auto px-6 py-3 mt-2" onPress={handleAddToCart}>
            <Text className="text-xl font-bold">AJOUTER AU PANIER</Text>
          </TouchableHighlight>
          <Text className="mx-auto text-3xl font-extrabold uppercase mt-4 mb-4">Produits Similaires</Text>
          {productEx.produitsSimilaires.map(productSingle => (
            <ProductCard key={`produitSimilaire-${productSingle.id}`} product={productSingle} />
          ))}
        </View>
      </View>
    </Page>
  )
}

export default ProductScreen
