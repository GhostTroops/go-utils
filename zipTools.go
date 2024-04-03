package go_utils

import (
	zip1 "archive/zip"
	"bytes"
	"log"
)

/*
	压缩多个文件

d 为 map[string][]byte{文件名1:文件内容1,文件名2:文件内容2,}

	var mJar = map[string][]byte{
			"buildServerResources/" + pluginName + ".jsp": X3Jsp}
		// pluginName+".zip"
		var mZip = map[string][]byte{
			"buildServerResources/" + pluginName + ".jsp": X3Jsp,
			"server/" + pluginName + ".jar":               CreateZip(mJar),
			"teamcity-plugin.xml": []byte(`?xml version="1.0" encoding="UTF-8"?>

<teamcity-plugin xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:noNamespaceSchemaLocation="urn:schemas-jetbrains-com:teamcity-plugin-v1-xml">

	<info>
	    <name>` + pluginName + `</name>
	    <display-name>` + pluginName + `</display-name>
	    <description>South rise such since land project enough.</description>
	    <version>1.0</version>
	    <vendor>
	        <name>Salas, Smith and Williams</name>
	        <url>http://meza-smith.biz/</url>
	    </vendor>
	</info>
	<deployment use-separate-classloader="true" node-responsibilities-aware="true"/>

</teamcity-plugin>`),

	}
	var data = CreateZip(mZip)
*/
func CreateZip(d map[string][]byte) []byte {
	buf := bytes.NewBuffer(nil)
	if zw := zip1.NewWriter(buf); nil != zw {
		for k, v := range d {
			// 创建文件写入器
			w, err := zw.Create(k)
			if err == nil { // 写入文件内容
				_, err = w.Write(v)
				if err != nil {
					log.Println(k, "Write", err)
				}
			} else {
				log.Println(k, "Create", err)
			}
		}
		zw.Close()
	}
	data := buf.Bytes()
	buf = nil
	return data
}
func GzipBytes(data []byte) []byte {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)
	if _, err := gz.Write(data); err != nil {
		return nil
	}
	if err := gz.Close(); err != nil {
		return nil
	}
	return buf.Bytes()
}

//func RmZipFile(zipFile string, rmFs ...string) {
//	jarWriter, err := zip33.OpenReader(zipFile)
//	if err != nil {
//		log.Println(err)
//		return
//	}
//	defer jarWriter.Close()
//	for _, file := range jarWriter.File {
//		for _, x := range rmFs {
//			if file.Name == x {
//				file.Remove()
//
//			}
//		}
//	}
//}
