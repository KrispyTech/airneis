import CategoryCard from "@/components/CategoryCard"
import Page from "@/components/Page"
import {getCategories} from "@/utils/services/categoryServices"
import {serverSideTranslations} from "next-i18next/serverSideTranslations"

export const getServerSideProps = async ({locale}) => {
    const {data: categories} = await getCategories()

    return {
        props: {
            categories,
            ...(await serverSideTranslations(locale, ["common"]))
        }
    }
}
const Home = props => {
    const {categories} = props

    return (
        <Page className="h-screen text-secondary">
            <div
                className="w-full h-1/4 sm:h-1/3 bg-[url('https://picsum.photos/1800/350')] bg-center flex items-center justify-center">
                <h1 className="text-white drop-shadow-textOutline font-bold text-xl lg:text-3xl">CATEGORY</h1>
            </div>
            <p className="font-semibold text-3xl px-8 py-14 text-center">FROM THE HIGHLANDS OF SCOTLAND OUR FURNITURE ARE IMMORTAL</p>
            <ul className="grid grid-cols-3 gap-x-12 gap-y-8 px-8 pb-6">
                {categories.map(category => (
                    <CategoryCard key={category.id} category={category}/>
                ))}
            </ul>
        </Page>
    )
}

export default Home
