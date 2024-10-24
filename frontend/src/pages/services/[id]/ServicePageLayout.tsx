import {LayoutProps} from "@/types/LayoutProps";
import {motion} from "framer-motion";
import animation from "@/utils/animation.json"

const ServicePageLayoutLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <motion.div
            {...animation}
        >
            {children}
        </motion.div>
    )
}

export default ServicePageLayoutLayout;