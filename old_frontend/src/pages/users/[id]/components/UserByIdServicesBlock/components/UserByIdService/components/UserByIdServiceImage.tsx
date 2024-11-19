import {getFile} from "@/utils/utils";
import {Image} from "@mantine/core";
import {UserService} from "@/types/UserService";

type UserByIdServiceImageProps = UserService;

const UserByIdServiceImage = (props: UserByIdServiceImageProps) => {
    const {image, title} = props;

    return (
        <Image
            className="object-cover aspect-video !rounded-md mb-3"
            src={getFile(image)}
            alt={title}
            height={200}
            width={200}
        />
    )
}

export default UserByIdServiceImage;