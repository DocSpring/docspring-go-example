# DocSpring Go Example

Here is an example Go program that will upload the `fw8ben.pdf` PDF template to your DocSpring account,
then list all templates.

# Usage

```
cd $GOPATH
cd src
git clone https://github.com/DocSpring/docspring-go-example.git
cd docspring-go-example

go run src/app/docspring.go --api_token "<API Token ID>:<API Token Secret>"
```
