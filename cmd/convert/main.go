package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/getkin/kin-openapi/openapi2"
	"github.com/getkin/kin-openapi/openapi2conv"
	"github.com/getkin/kin-openapi/openapi3"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
)

func main() {
	specsConvertedToV3DirName := "specs-v3"
	specsConvertedToV3Path := "./internal/" + specsConvertedToV3DirName
	swaggerModelsPath := "./internal/selling-partner-api-models/models"

	versionNumberRegex, err := regexp.Compile("(([^/]+)V[0-9]+).json")
	if err != nil {
		panic(err)
	}
	versionDateRegex, err := regexp.Compile("(([^/]+)_[0-9]{4}-[0-9]{2}-[0-9]{2}).json")
	if err != nil {
		panic(err)
	}
	noVersionRegex, err := regexp.Compile("([^/]+).json")
	if err != nil {
		panic(err)
	}
	if err = filepath.Walk(swaggerModelsPath, func(filePath string, info os.FileInfo, err error) error {
		if filepath.Ext(filePath) == ".json" {

			buf, err := ioutil.ReadFile(filePath)
			if err != nil {
				return err
			}

			var doc2 openapi2.T
			if err = json.Unmarshal(buf, &doc2); err != nil {
				return err
			}

			doc3, err := openapi2conv.ToV3(&doc2)
			if err != nil {
				return err
			}

			outBuf, err := doc3.MarshalJSON()
			if err != nil {
				return err
			}

			filename := path.Base(filePath)
			err = ioutil.WriteFile(filepath.Join(specsConvertedToV3Path, filename), outBuf, os.ModePerm)
			if err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		panic(err)
	}

	cmdStr := ""

	if err := filepath.Walk(specsConvertedToV3Path, func(filePath string, info os.FileInfo, err error) error {
		if filepath.Ext(filePath) == ".json" {

			buf, err := ioutil.ReadFile(filePath)
			if err != nil {
				return err
			}

			var doc3 openapi3.T
			if err = json.Unmarshal(buf, &doc3); err != nil {
				return err
			}

			packageName := ""
			versionedPath := ""
			if len(doc3.Paths) > 0 {
				for _, item := range doc3.Paths {
					if item.Get != nil && len(item.Get.Tags) > 0 {
						packageName = item.Get.Tags[0]
						break
					}
					if item.Post != nil && len(item.Post.Tags) > 0 {
						packageName = item.Post.Tags[0]
						break
					}
				}
				versionedPath = packageName
				foundMatches := versionNumberRegex.FindStringSubmatch(filePath)
				if foundMatches == nil {
					foundMatches = versionDateRegex.FindStringSubmatch(filePath)
				}
				if foundMatches == nil {
					foundMatches = noVersionRegex.FindStringSubmatch(filePath)
				}
				if foundMatches == nil {
					panic(errors.New("Failed to fetch version and package name for the file path " + filePath))
				}
				if len(foundMatches) == 2 {
					if len(packageName) == 0 {
						packageName = foundMatches[1]
					}
					versionedPath = foundMatches[1]
				} else if len(foundMatches) > 2 {
					if len(packageName) == 0 {
						packageName = foundMatches[2]
					}
					versionedPath = foundMatches[1]
				} else {
					panic(errors.New(fmt.Sprint("Too many matches ", foundMatches)))
				}
			}

			cmdStr += fmt.Sprintf("mkdir -p ./generated/%s/%s/ & ./gosdk-codegen -generate types --package=%s -o generated/%s/%s/types.gen.go internal/%s/%s \n", versionedPath, packageName, packageName, versionedPath, packageName, specsConvertedToV3DirName, path.Base(filePath))
			cmdStr += fmt.Sprintf("mkdir -p ./generated/%s/%s/ & ./gosdk-codegen -generate client --package=%s -o generated/%s/%s/api.gen.go internal/%s/%s \n", versionedPath, packageName, packageName, versionedPath, packageName, specsConvertedToV3DirName, path.Base(filePath))

		}
		return nil
	}); err != nil {
		panic(err)
	}

	_ = os.WriteFile("./gencode.sh", []byte(cmdStr), os.ModePerm)
}
