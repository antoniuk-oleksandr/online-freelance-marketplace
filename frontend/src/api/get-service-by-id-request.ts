import {Service} from "@/types/Service";
import {User} from "@/types/User";
import {Review} from "@/types/Review";

export const getServiceByIdRequest = async (id: string) => {
    const response = {
        id: 1,
        title: "Full-Stack Web Development",
        description: "I will build a complete full-stack web application using modern technologies like React, Node.js, and MongoDB.",
        images: ["https://wallpaperswide.com/download/nature_245-wallpaper-2560x1440.jpg"],
        createdAt: new Date("2024-10-21T00:00:00Z"),
        rating: 4.8,
        category: "Web Development",
        freelancer: {
            id: 1,
            avatar: "https://www.tjasakovac.com/wp-content/uploads/2023/02/LRM_EXPORT_124432641232861_20190823_173001741.jpg",
            firstName: "John",
            surname: "Doe",
            rating: 4.9,
            level: 3,
            reviewsCount: 150,
        } as User,
        reviews: [
            {
                "id": 1,
                "rating": 5,
                "content": "Outstanding service. Highly recommended!",
                "createdAt": "2024-10-13T22:08:40.184+00:00",
                "endedAt": "2024-10-14T22:08:40.184+00:00",
                "customer": {
                    "id": 3,
                    "firstName": "Jane",
                    "surname": "Doe",
                    "avatar": null
                },
                "service": {
                    "id": 3,
                    "price": 499.99,
                    "image": "image_1.jpg",
                    "title": "E-commerce Web Development"
                }
            } as unknown as Review,
            {
                "id": 4,
                "rating": 2,
                "content": "Not happy with the service. Needs improvement.",
                "createdAt": "2024-10-10T22:08:40.184+00:00",
                "endedAt": "2024-10-11T22:08:40.184+00:00",
                "customer": {
                    "id": 4,
                    "firstName": "Michael",
                    "surname": "Johnson",
                    "avatar": null
                },
                "service": {
                    "id": 4,
                    "price": 99.99,
                    "image": "image_1.jpg",
                    "title": "Logo Design"
                }
            } as unknown as Review,
            {
                "id": 2,
                "rating": 4,
                "content": "Very good service with minor delays.",
                "createdAt": "2024-10-05T22:08:40.184+00:00",
                "endedAt": "2024-10-06T22:08:40.184+00:00",
                "customer": {
                    "id": 6,
                    "firstName": "Chris",
                    "surname": "Brown",
                    "avatar": null
                },
                "service": {
                    "id": 6,
                    "price": 399.99,
                    "image": "image_1.jpg",
                    "title": "Graphic Design for Print"
                }
            } as unknown as Review,
            {
                "id": 5,
                "rating": 1,
                "content": "Very poor service. Would not recommend.",
                "createdAt": "2024-09-30T22:08:40.184+00:00",
                "endedAt": "2024-10-01T22:08:40.184+00:00",
                "customer": {
                    "id": 8,
                    "firstName": "David",
                    "surname": "Wilson",
                    "avatar": null
                },
                "service": {
                    "id": 1,
                    "price": 79.99,
                    "image": "image_1.jpg",
                    "title": "Custom Website Development"
                }
            } as unknown as Review
        ],
        packages: [
            {
                id: 1,
                title: "Basic Package",
                description: "A simple landing page with responsive design.",
                price: 500,
                deliveryDays: 5
            },
            {
                id: 2,
                title: "Standard Package",
                description: "A dynamic web application with backend integration.",
                price: 1500,
                deliveryDays: 10
            },
            {
                id: 3,
                title: "Premium Package",
                description: "A full-stack web application with all features including deployment and support.",
                price: 3000,
                deliveryDays: 15
            }
        ]
    } as Service

    return {
        data: response,
        status: 200
    }
}