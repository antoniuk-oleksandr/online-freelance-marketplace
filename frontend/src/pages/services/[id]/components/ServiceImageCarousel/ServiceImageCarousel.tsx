import {Carousel} from "@mantine/carousel";
import '@mantine/carousel/styles.css';
import {Service} from "@/types/Service";
import ServiceImageCarouselSlide from "../ServiceImageCarouselSlide/ServiceImageCarouselSlide";

type ServiceImageCarouselProps = Service;

const ServiceImageCarousel = (props: ServiceImageCarouselProps) => {
    const {images} = props;

    return (
        <Carousel loop withIndicators>
            {images.map((item, index) => (
                <ServiceImageCarouselSlide key={index} src={item}/>
            ))}
        </Carousel>
    )
}

export default ServiceImageCarousel;