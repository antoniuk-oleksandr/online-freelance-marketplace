import RightSignSideLayout from "./RightSignSideLayout";
import Image from "next/image";

const RightSignSide = () => {
    return (
        <RightSignSideLayout>
            <Image
                draggable={false}
                className={"object-center object-cover size-full select-none"}
                width={1024}
                height={1024}
                src={'/images/sign-form.webp'}
                alt={'image'}
            />
        </RightSignSideLayout>
    )
}

export default RightSignSide;