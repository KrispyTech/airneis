import {Button, FlatList, ScrollView, Text} from "react-native"
import Page from "../components/Page"
import screens from "../utils/constants/screens"
import {getCategories} from "../utils/categoriesServices"
import {useEffect, useState} from "react"
import ErrorScreen from "./ErrorScreen"
import errors from "../utils/constants/errors"

const HomeScreen = props => {
    const {navigation} = props
    const [categories, setCategories] = useState([])
    const [hadError, setHadError] = useState(false)
    const goToCategory = id => () => navigation.navigate(screens.categoryPage, {id})

    useEffect(() => {
        ;(async () => {
            try {
                const categoriesResult = await getCategories()
                setCategories(categoriesResult.data)
            } catch (e) {
                setHadError(true)
            }
        })()
    }, [])

    if (hadError) {
        return <ErrorScreen navigation={navigation} text={errors.homePage.unableToLoadCategories}/>
    }

    return (
        <Page navigation={navigation}>
            <ScrollView>
                <Text>Home page</Text>
                {categories.map(({name, ID: id}) => (
                    <Button title={name} onPress={goToCategory(id)} key={id}/>
                ))}
            </ScrollView>
        </Page>
    )
}

export default HomeScreen
