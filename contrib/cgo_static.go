//go:build !customenv && static
// +build !customenv,static

package contrib

// Changes here should be mirrored in contrib/cgo_static.go and cuda/cgo_static.go.

/*
#cgo CXXFLAGS:   --std=c++11
#cgo !windows CPPFLAGS: -I/usr/local/include -I/usr/local/include/opencv4
#cgo !windows LDFLAGS: -L/usr/local/lib -L/usr/local/lib/opencv4/3rdparty -lopencv_gapi -lopencv_stitching -lopencv_aruco -lopencv_bgsegm -lopencv_bioinspired -lopencv_ccalib -lopencv_dnn_objdetect -lopencv_dpm -lopencv_face -lopencv_fuzzy -lopencv_hfs -lopencv_img_hash -lopencv_line_descriptor -lopencv_quality -lopencv_reg -lopencv_rgbd -lopencv_saliency -lopencv_stereo -lopencv_structured_light -lopencv_phase_unwrapping -lopencv_superres -lopencv_optflow -lopencv_surface_matching -lopencv_tracking -lopencv_datasets -lopencv_text -lopencv_highgui -lopencv_dnn -lopencv_plot -lopencv_videostab -lopencv_video -lopencv_videoio -lopencv_xfeatures2d -lopencv_shape -lopencv_ml -lopencv_ximgproc -lopencv_xobjdetect -lopencv_objdetect -lopencv_calib3d -lopencv_imgcodecs -lopencv_features2d -lopencv_flann -lopencv_xphoto -lopencv_wechat_qrcode -lopencv_photo -lopencv_imgproc -lopencv_core -littnotify -llibprotobuf -lIlmImf -lquirc -lippiw -lippicv -lade -lz -ljpeg -ldl -lm -lpthread -lrt -lquadmath
#cgo windows  CPPFLAGS:   -IC:/opencv/build/install/include
#cgo windows  LDFLAGS:    -LC:/opencv/build/install/x64/mingw/staticlib -lopencv_stereo410 -lopencv_tracking410 -lopencv_superres410 -lopencv_stitching410 -lopencv_optflow410 -lopencv_gapi410 -lopencv_face410 -lopencv_dpm410 -lopencv_dnn_objdetect410 -lopencv_ccalib410 -lopencv_bioinspired410 -lopencv_bgsegm410 -lopencv_aruco410 -lopencv_xobjdetect410 -lopencv_ximgproc410 -lopencv_xfeatures2d410 -lopencv_videostab410 -lopencv_video410 -lopencv_structured_light410 -lopencv_shape410 -lopencv_rgbd410 -lopencv_rapid410 -lopencv_objdetect410 -lopencv_mcc410 -lopencv_highgui410 -lopencv_datasets410 -lopencv_calib3d410 -lopencv_videoio410 -lopencv_text410 -lopencv_line_descriptor410 -lopencv_imgcodecs410 -lopencv_img_hash410 -lopencv_hfs410 -lopencv_fuzzy410 -lopencv_features2d410 -lopencv_dnn_superres410 -lopencv_dnn410 -lopencv_xphoto410 -lopencv_wechat_qrcode410 -lopencv_surface_matching410 -lopencv_reg410 -lopencv_quality410 -lopencv_plot410 -lopencv_photo410 -lopencv_phase_unwrapping410 -lopencv_ml410 -lopencv_intensity_transform410 -lopencv_imgproc410 -lopencv_flann410 -lopencv_core410 -lade -lquirc -llibprotobuf -lIlmImf -llibpng -llibopenjp2 -llibwebp -llibtiff -llibjpeg-turbo -lzlib -lkernel32 -lgdi32 -lwinspool -lshell32 -lole32 -loleaut32 -luuid -lcomdlg32 -ladvapi32 -luser32
*/
import "C"
