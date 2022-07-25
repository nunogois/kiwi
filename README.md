# 🥝 Kiwi

Kiwi is an extremely simple and performant key-value store written in Go.

## Routes

Kiwi follows a somewhat standard REST API pattern.

 - **GET** `/store/my/custom/id`
    - Returns data in the `my/custom/id` key
 - **POST** `/store/my/custom/id`
    - Stores (replacing) data in the `my/custom/id` key
 - **PATCH** `/store/my/custom/id`
    - Patches data in the `my/custom/id` key
 - **DELETE** `/store/my/custom/id`
    - Deletes data in the `my/custom/id` key
