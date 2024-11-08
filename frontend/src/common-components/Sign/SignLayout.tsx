import {LayoutProps} from "@/types/LayoutProps";
import animation from "@/utils/animation.json";
import PaperElement from "@/common-components/PaperElement";
import {motion} from "framer-motion";

const SignLayoutLayout = (props: LayoutProps) => {
    const {children} = props;

    return (
        <motion.div
            className={"min-h-svh flex items-center bg-light-palette-background-block dark:bg-dark-palette-background-block lg:bg-transparent dark:lg:bg-transparent -mx-4 lg:mx-0 justify-center flex-1 flex-grow"}
            {...animation}
        >
            <PaperElement
                styles={"w-full min-h-212 rounded-none lg:rounded-lg !bg-light-palette-background-block dark:!bg-dark-palette-background-block !p-0 grid grid-cols-1 lg:grid-cols-2 overflow-hidden border-none lg:border-solid bg-transparent"}
            >
                {children}
            </PaperElement>
        </motion.div>
    )
}

export default SignLayoutLayout;