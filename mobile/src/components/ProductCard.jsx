import { Text, View, Image } from "react-native"
import mobileConfig from "../../mobileConfig"

const ProductCard = props => {
  const {
    product: { name, thumbnailUrl, priceWithTaxes }
  } = props

  return (
    <View>
      <Image
        style={mobileConfig.images.fullScreenImageMd}
        className="rounded"
        source={{
          uri: thumbnailUrl
        }}
      />
      <View className="flex flex-row justify-between py-1">
        <Text className="text-lg font-light">{name} </Text>
        <Text className="text-lg font-light">{priceWithTaxes} €</Text>
      </View>
    </View>
  )
}

export default ProductCard
