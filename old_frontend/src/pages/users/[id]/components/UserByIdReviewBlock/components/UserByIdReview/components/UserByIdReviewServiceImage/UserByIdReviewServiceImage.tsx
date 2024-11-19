import UserByIdReviewServiceImageLayout from "./UserByIdReviewServiceImageLayout";
import {Image} from "@mantine/core";
import {getFile} from "@/utils/utils";

type UserByIdReviewServiceImageProps = {
    src: string,
}

const UserByIdReviewServiceImage = (props: UserByIdReviewServiceImageProps) => {
    const {src} = props;

    return (
        <UserByIdReviewServiceImageLayout>
            <Image
                className={"object-center aspect-video w-full object-cover h-9"}
                src={getFile(src)}
                alt={"service"}
                width={48}
                height={48}
            />
        </UserByIdReviewServiceImageLayout>
    )
}

export default UserByIdReviewServiceImage;