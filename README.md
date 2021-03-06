# Ires
`Ires` is image resizer gem.

## Usage

```erb
<!-- Usually -->
<%= ires_tag( path: "image_01.jpg", width: 90, height: 120, mode: "resize" ) %>

<!-- Using image_tag options -->
<%= ires_tag( path: "http://example.com/image_02.jpg", width: 200, height: 200, mode: "crop", alt: "example image" ) %>
```

### Select mode

| info                       |     　　　mode 　　　  |
|:---------------------------|:--------------------:|
| Resize                     | resize               |
| Cropping                   | crop                 |
| Rsize after Cropping       | rsize_to_crop        | 

### Saved directory

```
.
└── public
    └── ires
        ├── crop
        │   └── 300x220
        │       └── image_300x220.jpg
        ├── original
        │   └── original
        │       └── image.jpg
        ├── resize
        │   ├── 200x150
        │   │   └── image_200x150.jpg
        │   ├── 300x220
        │   │   └── image_300x220.jpg
        │   ├── 300x400
        │   │   └── image_300x400.jpg
        │   ├── 400x300
        │   │   └── image_400x300.jpg
        │   └── 90x120
        │       └── image_90x120.jpg
        └── resize_to_crop
            └── 300x220
                └── image_300x220.jpg
```

`original` directory where downloaded images are saved.

## Installation
Add this line to your application's Gemfile:

```ruby
gem 'ires'
```

And then execute:
```bash
$ bundle
```

Or install it yourself as:
```bash
$ gem install ires
```

## Development

環境はDockerで準備しています

```shell
$ docker-compose up

# コンテナに入る
# go
$ docker-compose exec ext bash
# ruby
$ docker-compose exec gem bash
```

## Gemテスト

### 1. Go（shared objectの作成）

パッケージ管理は[dep](https://github.com/golang/dep)を使っています

```shell
$ docker-compose exec ext bash
$ dep ensure

# shared object として出力する
$ go build -buildmode=c-shared -o shared/ires.so main.go 
```

### 2. Railsアプリの起動

```shell
$ docker-compose exec gem bash
$ cd test/dummy
$ bin/rails s -b 0.0.0.0
```

## License
The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
