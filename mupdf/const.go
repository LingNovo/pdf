// THE AUTOGENERATED LICENSE. ALL THE RIGHTS ARE RESERVED BY ROBOTS.

// WARNING: This file has automatically been generated on Wed, 07 Mar 2018 17:31:54 CST.
// By https://git.io/c-for-go. DO NOT EDIT.

package mupdf

/*
#include "mupdf/pdf.h"
#include "mupdf/fitz.h"
#include <stdlib.h>
#include "cgo_helpers.h"
*/
import "C"

const (
	// FzVersion as defined in fitz/version.h:4
	FzVersion = "1.12.0"
	// FzPlottersG as defined in fitz/config.h:98
	FzPlottersG = 1
	// FzPlottersRgb as defined in fitz/config.h:102
	FzPlottersRgb = 1
	// FzPlottersCmyk as defined in fitz/config.h:106
	FzPlottersCmyk = 1
	// FzPlottersN as defined in fitz/config.h:110
	FzPlottersN = 1
	// FzEnablePDF as defined in fitz/config.h:120
	FzEnablePDF = 1
	// FzEnableXps as defined in fitz/config.h:124
	FzEnableXps = 1
	// FzEnableSvg as defined in fitz/config.h:128
	FzEnableSvg = 1
	// FzEnableCbz as defined in fitz/config.h:132
	FzEnableCbz = 1
	// FzEnableImg as defined in fitz/config.h:136
	FzEnableImg = 1
	// FzEnableTiff as defined in fitz/config.h:140
	FzEnableTiff = 1
	// FzEnableHtml as defined in fitz/config.h:144
	FzEnableHtml = 1
	// FzEnableEpub as defined in fitz/config.h:148
	FzEnableEpub = 1
	// FzEnableGprf as defined in fitz/config.h:152
	FzEnableGprf = 0
	// FzEnableJpx as defined in fitz/config.h:156
	FzEnableJpx = 1
	// FzEnableJs as defined in fitz/config.h:160
	FzEnableJs = 1
	// FzPi as defined in fitz/system.h:37
	FzPi = 3.1415927
	// FzRadian as defined in fitz/system.h:38
	FzRadian = 57.29578
	// FzDegree as defined in fitz/system.h:39
	FzDegree = 0.017453292
	// FzSqrt2 as defined in fitz/system.h:40
	FzSqrt2 = 1.4142135
	// FzLn2 as defined in fitz/system.h:41
	FzLn2 = 0.6931472
	// FzJmpBuf as defined in fitz/system.h:78
	FzJmpBuf = 0
	// FzAesDecrypt as defined in fitz/crypt.h:105
	FzAesDecrypt = 0
	// FzAesEncrypt as defined in fitz/crypt.h:106
	FzAesEncrypt = 1
	// FzReplacementCharacter as defined in fitz/string-util.h:7
	FzReplacementCharacter = 65533
	// FzMetaFormat as defined in fitz/document.h:598
	FzMetaFormat = "format"
	// FzMetaEncryption as defined in fitz/document.h:599
	FzMetaEncryption = "encryption"
	// FzMetaInfoAuthor as defined in fitz/document.h:601
	FzMetaInfoAuthor = "info:Author"
	// FzMetaInfoTitle as defined in fitz/document.h:602
	FzMetaInfoTitle = "info:Title"
	// PDFMrangeCap as defined in pdf/cmap.h:13
	PDFMrangeCap = 8
)

// FzBidiDirection as declared in fitz/bidi.h:31
type FzBidiDirection int32

// FzBidiDirection enumeration from fitz/bidi.h:31
const (
	FzBidiLtr     FzBidiDirection = iota
	FzBidiRtl     FzBidiDirection = 1
	FzBidiNeutral FzBidiDirection = 2
)

// FzBidiFlags as declared in fitz/bidi.h:38
type FzBidiFlags int32

// FzBidiFlags enumeration from fitz/bidi.h:38
const (
	FzBidiClassifyWhiteSpace FzBidiFlags = 1
	FzBidiReplaceTab         FzBidiFlags = 2
)

// FzLinecap as declared in fitz/path.h:26
type FzLinecap int32

// FzLinecap enumeration from fitz/path.h:26
const (
	FzLinecapButt     FzLinecap = iota
	FzLinecapRound    FzLinecap = 1
	FzLinecapSquare   FzLinecap = 2
	FzLinecapTriangle FzLinecap = 3
)

// FzLinejoin as declared in fitz/path.h:34
type FzLinejoin int32

// FzLinejoin enumeration from fitz/path.h:34
const (
	FzLinejoinMiter    FzLinejoin = iota
	FzLinejoinRound    FzLinejoin = 1
	FzLinejoinBevel    FzLinejoin = 2
	FzLinejoinMiterXps FzLinejoin = 3
)

// FzTextLanguage as declared in fitz/text.h:46
type FzTextLanguage int32

// FzTextLanguage enumeration from fitz/text.h:46
const (
	FzLangUnset  FzTextLanguage = iota
	FzLangUr     FzTextLanguage = 507
	FzLangUrd    FzTextLanguage = 3423
	FzLangKo     FzTextLanguage = 416
	FzLangJa     FzTextLanguage = 37
	FzLangZh     FzTextLanguage = 242
	FzLangZhHans FzTextLanguage = 14093
	FzLangZhHant FzTextLanguage = 14822
)

// FzDeflateLevel as declared in fitz/compress.h:12
type FzDeflateLevel int32

// FzDeflateLevel enumeration from fitz/compress.h:12
const (
	FzDeflateNone      FzDeflateLevel = iota
	FzDeflateBestSpeed FzDeflateLevel = 1
	FzDeflateBest      FzDeflateLevel = 9
	FzDeflateDefault   FzDeflateLevel = -1
)

// FzSeparationBehavior as declared in fitz/separation.h:32
type FzSeparationBehavior int32

// FzSeparationBehavior enumeration from fitz/separation.h:32
const (
	FzSeparationComposite FzSeparationBehavior = iota
	FzSeparationSpot      FzSeparationBehavior = 1
	FzSeparationDisabled  FzSeparationBehavior = 2
)

// FzPermission as declared in fitz/document.h:29
type FzPermission int32

// FzPermission enumeration from fitz/document.h:29
const (
	FzPermissionPrint    FzPermission = 112
	FzPermissionCopy     FzPermission = 99
	FzPermissionEdit     FzPermission = 101
	FzPermissionAnnotate FzPermission = 110
)

// PDFLayerConfigUiType as declared in pdf/document.h:231
type PDFLayerConfigUiType int32

// PDFLayerConfigUiType enumeration from pdf/document.h:231
const (
	PDFLayerUiLabel    PDFLayerConfigUiType = iota
	PDFLayerUiCheckbox PDFLayerConfigUiType = 1
	PDFLayerUiRadiobox PDFLayerConfigUiType = 2
)

// PDFPortfolioSchemaType as declared in pdf/document.h:297
type PDFPortfolioSchemaType int32

// PDFPortfolioSchemaType enumeration from pdf/document.h:297
const (
	PDFSchemaNumber       PDFPortfolioSchemaType = iota
	PDFSchemaSize         PDFPortfolioSchemaType = 1
	PDFSchemaText         PDFPortfolioSchemaType = 2
	PDFSchemaDate         PDFPortfolioSchemaType = 3
	PDFSchemaDesc         PDFPortfolioSchemaType = 4
	PDFSchemaModdate      PDFPortfolioSchemaType = 5
	PDFSchemaCreationdate PDFPortfolioSchemaType = 6
	PDFSchemaFilename     PDFPortfolioSchemaType = 7
	PDFSchemaUnknown      PDFPortfolioSchemaType = 8
)

// PDFToken as declared in pdf/parse.h:20
type PDFToken int32

// PDFToken enumeration from pdf/parse.h:20
const (
	PDFTokError      PDFToken = iota
	PDFTokEof        PDFToken = 1
	PDFTokOpenArray  PDFToken = 2
	PDFTokCloseArray PDFToken = 3
	PDFTokOpenDict   PDFToken = 4
	PDFTokCloseDict  PDFToken = 5
	PDFTokOpenBrace  PDFToken = 6
	PDFTokCloseBrace PDFToken = 7
	PDFTokName       PDFToken = 8
	PDFTokInt        PDFToken = 9
	PDFTokReal       PDFToken = 10
	PDFTokString     PDFToken = 11
	PDFTokKeyword    PDFToken = 12
	PDFTokR          PDFToken = 13
	PDFTokTrue       PDFToken = 14
	PDFTokFalse      PDFToken = 15
	PDFTokNull       PDFToken = 16
	PDFTokObj        PDFToken = 17
	PDFTokEndobj     PDFToken = 18
	PDFTokStream     PDFToken = 19
	PDFTokEndstream  PDFToken = 20
	PDFTokXref       PDFToken = 21
	PDFTokTrailer    PDFToken = 22
	PDFTokStartxref  PDFToken = 23
	PDFNumTokens     PDFToken = 24
)

// FzAnnotType as declared in pdf/annot.h:32
type FzAnnotType int32

// FzAnnotType enumeration from pdf/annot.h:32
const (
	PDFAnnotText           FzAnnotType = iota
	PDFAnnotLink           FzAnnotType = 1
	PDFAnnotFreeText       FzAnnotType = 2
	PDFAnnotLine           FzAnnotType = 3
	PDFAnnotSquare         FzAnnotType = 4
	PDFAnnotCircle         FzAnnotType = 5
	PDFAnnotPolygon        FzAnnotType = 6
	PDFAnnotPolyLine       FzAnnotType = 7
	PDFAnnotHighlight      FzAnnotType = 8
	PDFAnnotUnderline      FzAnnotType = 9
	PDFAnnotSquiggly       FzAnnotType = 10
	PDFAnnotStrikeOut      FzAnnotType = 11
	PDFAnnotStamp          FzAnnotType = 12
	PDFAnnotCaret          FzAnnotType = 13
	PDFAnnotInk            FzAnnotType = 14
	PDFAnnotPopup          FzAnnotType = 15
	PDFAnnotFileAttachment FzAnnotType = 16
	PDFAnnotSound          FzAnnotType = 17
	PDFAnnotMovie          FzAnnotType = 18
	PDFAnnotWidget         FzAnnotType = 19
	PDFAnnotScreen         FzAnnotType = 20
	PDFAnnotPrinterMark    FzAnnotType = 21
	PDFAnnotTrapNet        FzAnnotType = 22
	PDFAnnotWatermark      FzAnnotType = 23
	PDFAnnot3d             FzAnnotType = 24
	PDFAnnotUnknown        FzAnnotType = -1
)

const (
	// FzErrorNone as declared in fitz/context.h:91
	FzErrorNone = iota
	// FzErrorMemory as declared in fitz/context.h:92
	FzErrorMemory = 1
	// FzErrorGeneric as declared in fitz/context.h:93
	FzErrorGeneric = 2
	// FzErrorSyntax as declared in fitz/context.h:94
	FzErrorSyntax = 3
	// FzErrorTrylater as declared in fitz/context.h:95
	FzErrorTrylater = 4
	// FzErrorAbort as declared in fitz/context.h:96
	FzErrorAbort = 5
	// FzErrorCount as declared in fitz/context.h:97
	FzErrorCount = 6
)

const (
	// FzLockAlloc as declared in fitz/context.h:140
	FzLockAlloc = iota
	// FzLockFreetype as declared in fitz/context.h:141
	FzLockFreetype = 1
	// FzLockGlyphcache as declared in fitz/context.h:142
	FzLockGlyphcache = 2
	// FzLockMax as declared in fitz/context.h:143
	FzLockMax = 3
)

const (
	// FzStoreUnlimited as declared in fitz/context.h:176
	FzStoreUnlimited = iota
	// FzStoreDefault as declared in fitz/context.h:177
	FzStoreDefault = 268435456
)

const (
	// FzUtfmax as declared in fitz/string-util.h:93
	FzUtfmax = 4
)

const (
	// FzStreamMetaProgressive as declared in fitz/stream.h:188
	FzStreamMetaProgressive = 1
	// FzStreamMetaLength as declared in fitz/stream.h:189
	FzStreamMetaLength = 2
)

const (
	// FzImageUnknown as declared in fitz/compressed-buffer.h:22
	FzImageUnknown = iota
	// FzImageRaw as declared in fitz/compressed-buffer.h:25
	FzImageRaw = 1
	// FzImageFax as declared in fitz/compressed-buffer.h:28
	FzImageFax = 2
	// FzImageFlate as declared in fitz/compressed-buffer.h:29
	FzImageFlate = 3
	// FzImageLzw as declared in fitz/compressed-buffer.h:30
	FzImageLzw = 4
	// FzImageRld as declared in fitz/compressed-buffer.h:31
	FzImageRld = 5
	// FzImageBmp as declared in fitz/compressed-buffer.h:34
	FzImageBmp = 6
	// FzImageGif as declared in fitz/compressed-buffer.h:35
	FzImageGif = 7
	// FzImageJpeg as declared in fitz/compressed-buffer.h:36
	FzImageJpeg = 8
	// FzImageJpx as declared in fitz/compressed-buffer.h:37
	FzImageJpx = 9
	// FzImageJxr as declared in fitz/compressed-buffer.h:38
	FzImageJxr = 10
	// FzImagePng as declared in fitz/compressed-buffer.h:39
	FzImagePng = 11
	// FzImagePnm as declared in fitz/compressed-buffer.h:40
	FzImagePnm = 12
	// FzImageTiff as declared in fitz/compressed-buffer.h:41
	FzImageTiff = 13
)

const (
	// FzMaxColors as declared in fitz/colorspace.h:8
	FzMaxColors = 32
)

const (
	// FzRiPerceptual as declared in fitz/colorspace.h:13
	FzRiPerceptual = iota
	// FzRiRelativeColorimetric as declared in fitz/colorspace.h:14
	FzRiRelativeColorimetric = 1
	// FzRiSaturation as declared in fitz/colorspace.h:15
	FzRiSaturation = 2
	// FzRiAbsoluteColorimetric as declared in fitz/colorspace.h:16
	FzRiAbsoluteColorimetric = 3
)

const (
	// FzCsDeviceGray as declared in fitz/colorspace.h:131
	FzCsDeviceGray = 1
	// FzCsDeviceN as declared in fitz/colorspace.h:132
	FzCsDeviceN = 2
	// FzCsSubtractive as declared in fitz/colorspace.h:133
	FzCsSubtractive = 4
	// FzCsLastPublicFlag as declared in fitz/colorspace.h:135
	FzCsLastPublicFlag = 4
)

const (
	// FzPixmapFlagInterpolate as declared in fitz/pixmap.h:370
	FzPixmapFlagInterpolate = 1
	// FzPixmapFlagFreeSamples as declared in fitz/pixmap.h:371
	FzPixmapFlagFreeSamples = 2
)

const (
	// FzMaxSeparations as declared in fitz/separation.h:17
	FzMaxSeparations = 64
)

const (
	// FzFunctionBased as declared in fitz/shade.h:18
	FzFunctionBased = 1
	// FzLinear as declared in fitz/shade.h:19
	FzLinear = 2
	// FzRadial as declared in fitz/shade.h:20
	FzRadial = 3
	// FzMeshType4 as declared in fitz/shade.h:21
	FzMeshType4 = 4
	// FzMeshType5 as declared in fitz/shade.h:22
	FzMeshType5 = 5
	// FzMeshType6 as declared in fitz/shade.h:23
	FzMeshType6 = 6
	// FzMeshType7 as declared in fitz/shade.h:24
	FzMeshType7 = 7
)

const (
	// FzAdobeCns1 as declared in fitz/font.h:58
	FzAdobeCns1 = iota
	// FzAdobeGb1 as declared in fitz/font.h:58
	FzAdobeGb1 = 1
	// FzAdobeJapan1 as declared in fitz/font.h:58
	FzAdobeJapan1 = 2
	// FzAdobeKorea1 as declared in fitz/font.h:58
	FzAdobeKorea1 = 3
)

const (
	// FzDevflagMask as declared in fitz/device.h:29
	FzDevflagMask = 1
	// FzDevflagColor as declared in fitz/device.h:30
	FzDevflagColor = 2
	// FzDevflagUncacheable as declared in fitz/device.h:31
	FzDevflagUncacheable = 4
	// FzDevflagFillcolorUndefined as declared in fitz/device.h:32
	FzDevflagFillcolorUndefined = 8
	// FzDevflagStrokecolorUndefined as declared in fitz/device.h:33
	FzDevflagStrokecolorUndefined = 16
	// FzDevflagStartcapUndefined as declared in fitz/device.h:34
	FzDevflagStartcapUndefined = 32
	// FzDevflagDashcapUndefined as declared in fitz/device.h:35
	FzDevflagDashcapUndefined = 64
	// FzDevflagEndcapUndefined as declared in fitz/device.h:36
	FzDevflagEndcapUndefined = 128
	// FzDevflagLinejoinUndefined as declared in fitz/device.h:37
	FzDevflagLinejoinUndefined = 256
	// FzDevflagMiterlimitUndefined as declared in fitz/device.h:38
	FzDevflagMiterlimitUndefined = 512
	// FzDevflagLinewidthUndefined as declared in fitz/device.h:39
	FzDevflagLinewidthUndefined = 1024
	// FzDevflagBboxDefined as declared in fitz/device.h:43
	FzDevflagBboxDefined = 2048
	// FzDevflagGridfitAsTiled as declared in fitz/device.h:44
	FzDevflagGridfitAsTiled = 4096
)

const (
	// FzBlendNormal as declared in fitz/device.h:50
	FzBlendNormal = iota
	// FzBlendMultiply as declared in fitz/device.h:51
	FzBlendMultiply = 1
	// FzBlendScreen as declared in fitz/device.h:52
	FzBlendScreen = 2
	// FzBlendOverlay as declared in fitz/device.h:53
	FzBlendOverlay = 3
	// FzBlendDarken as declared in fitz/device.h:54
	FzBlendDarken = 4
	// FzBlendLighten as declared in fitz/device.h:55
	FzBlendLighten = 5
	// FzBlendColorDodge as declared in fitz/device.h:56
	FzBlendColorDodge = 6
	// FzBlendColorBurn as declared in fitz/device.h:57
	FzBlendColorBurn = 7
	// FzBlendHardLight as declared in fitz/device.h:58
	FzBlendHardLight = 8
	// FzBlendSoftLight as declared in fitz/device.h:59
	FzBlendSoftLight = 9
	// FzBlendDifference as declared in fitz/device.h:60
	FzBlendDifference = 10
	// FzBlendExclusion as declared in fitz/device.h:61
	FzBlendExclusion = 11
	// FzBlendHue as declared in fitz/device.h:64
	FzBlendHue = 12
	// FzBlendSaturation as declared in fitz/device.h:65
	FzBlendSaturation = 13
	// FzBlendColor as declared in fitz/device.h:66
	FzBlendColor = 14
	// FzBlendLuminosity as declared in fitz/device.h:67
	FzBlendLuminosity = 15
	// FzBlendModemask as declared in fitz/device.h:70
	FzBlendModemask = 15
	// FzBlendIsolated as declared in fitz/device.h:71
	FzBlendIsolated = 16
	// FzBlendKnockout as declared in fitz/device.h:72
	FzBlendKnockout = 32
)

const (
	// FzDontInterpolateImages as declared in fitz/device.h:193
	FzDontInterpolateImages = 1
	// FzMaintainContainerStack as declared in fitz/device.h:194
	FzMaintainContainerStack = 2
	// FzNoCache as declared in fitz/device.h:195
	FzNoCache = 4
)

const (
	// FzTestOptImages as declared in fitz/device.h:306
	FzTestOptImages = 1
	// FzTestOptShadings as declared in fitz/device.h:310
	FzTestOptShadings = 2
)

// FzDrawOptionsUsage as declared in fitz/device.h:396
var FzDrawOptionsUsage string

const (
	// FzStextPreserveLigatures as declared in fitz/structured-text.h:43
	FzStextPreserveLigatures = 1
	// FzStextPreserveWhitespace as declared in fitz/structured-text.h:44
	FzStextPreserveWhitespace = 2
	// FzStextPreserveImages as declared in fitz/structured-text.h:45
	FzStextPreserveImages = 4
)

const (
	// FzStextBlockText as declared in fitz/structured-text.h:60
	FzStextBlockText = iota
	// FzStextBlockImage as declared in fitz/structured-text.h:61
	FzStextBlockImage = 1
)

// FzStextOptionsUsage as declared in fitz/structured-text.h:104
var FzStextOptionsUsage string

const (
	// FzTransitionNone as declared in fitz/transition.h:11
	FzTransitionNone = iota
	// FzTransitionSplit as declared in fitz/transition.h:12
	FzTransitionSplit = 1
	// FzTransitionBlinds as declared in fitz/transition.h:13
	FzTransitionBlinds = 2
	// FzTransitionBox as declared in fitz/transition.h:14
	FzTransitionBox = 3
	// FzTransitionWipe as declared in fitz/transition.h:15
	FzTransitionWipe = 4
	// FzTransitionDissolve as declared in fitz/transition.h:16
	FzTransitionDissolve = 5
	// FzTransitionGlitter as declared in fitz/transition.h:17
	FzTransitionGlitter = 6
	// FzTransitionFly as declared in fitz/transition.h:18
	FzTransitionFly = 7
	// FzTransitionPush as declared in fitz/transition.h:19
	FzTransitionPush = 8
	// FzTransitionCover as declared in fitz/transition.h:20
	FzTransitionCover = 9
	// FzTransitionUncover as declared in fitz/transition.h:21
	FzTransitionUncover = 10
	// FzTransitionFade as declared in fitz/transition.h:22
	FzTransitionFade = 11
)

// FzPDFWriteOptionsUsage as declared in fitz/writer.h:152
var FzPDFWriteOptionsUsage string

// FzSvgWriteOptionsUsage as declared in fitz/writer.h:153
var FzSvgWriteOptionsUsage string

// FzPclWriteOptionsUsage as declared in fitz/writer.h:155
var FzPclWriteOptionsUsage string

// FzPclmWriteOptionsUsage as declared in fitz/writer.h:156
var FzPclmWriteOptionsUsage string

// FzPwgWriteOptionsUsage as declared in fitz/writer.h:157
var FzPwgWriteOptionsUsage string

const (
	// FzSvgTextAsPath as declared in fitz/output-svg.h:10
	FzSvgTextAsPath = iota
	// FzSvgTextAsText as declared in fitz/output-svg.h:11
	FzSvgTextAsText = 1
)

const ()

const (
	// PDFFlagsMemoBm as declared in pdf/object.h:72
	PDFFlagsMemoBm = iota
	// PDFFlagsMemoOp as declared in pdf/object.h:73
	PDFFlagsMemoOp = 1
)

const (
	// PDFLEXBUFSMALLSIZE as declared in pdf/document.h:19
	PDFLEXBUFSMALLSIZE = 256
	// PDFLEXBUFLARGESIZE as declared in pdf/document.h:20
	PDFLEXBUFLARGESIZE = 65536
)

const ()

const (
	// PDFPageIncompleteContents as declared in pdf/page.h:240
	PDFPageIncompleteContents = 1
	// PDFPageIncompleteAnnots as declared in pdf/page.h:241
	PDFPageIncompleteAnnots = 2
)

const (
	// PDFFdFixedPitch as declared in pdf/font.h:12
	PDFFdFixedPitch = 1
	// PDFFdSerif as declared in pdf/font.h:13
	PDFFdSerif = 2
	// PDFFdSymbolic as declared in pdf/font.h:14
	PDFFdSymbolic = 4
	// PDFFdScript as declared in pdf/font.h:15
	PDFFdScript = 8
	// PDFFdNonsymbolic as declared in pdf/font.h:16
	PDFFdNonsymbolic = 32
	// PDFFdItalic as declared in pdf/font.h:17
	PDFFdItalic = 64
	// PDFFdAllCap as declared in pdf/font.h:18
	PDFFdAllCap = 65536
	// PDFFdSmallCap as declared in pdf/font.h:19
	PDFFdSmallCap = 131072
	// PDFFdForceBold as declared in pdf/font.h:20
	PDFFdForceBold = 262144
)

// PDFDocEncoding as declared in pdf/font.h:27
var PDFDocEncoding [256]uint16

// PDFMacRoman as declared in pdf/font.h:28
var PDFMacRoman [256]string

// PDFMacExpert as declared in pdf/font.h:29
var PDFMacExpert [256]string

// PDFWinAnsi as declared in pdf/font.h:30
var PDFWinAnsi [256]string

// PDFStandard as declared in pdf/font.h:31
var PDFStandard [256]string

const (
	// PDFAnnotIsInvisible as declared in pdf/annot.h:39
	PDFAnnotIsInvisible = 1
	// PDFAnnotIsHidden as declared in pdf/annot.h:40
	PDFAnnotIsHidden = 2
	// PDFAnnotIsPrint as declared in pdf/annot.h:41
	PDFAnnotIsPrint = 4
	// PDFAnnotIsNoZoom as declared in pdf/annot.h:42
	PDFAnnotIsNoZoom = 8
	// PDFAnnotIsNoRotate as declared in pdf/annot.h:43
	PDFAnnotIsNoRotate = 16
	// PDFAnnotIsNoView as declared in pdf/annot.h:44
	PDFAnnotIsNoView = 32
	// PDFAnnotIsReadOnly as declared in pdf/annot.h:45
	PDFAnnotIsReadOnly = 64
	// PDFAnnotIsLocked as declared in pdf/annot.h:46
	PDFAnnotIsLocked = 128
	// PDFAnnotIsToggleNoView as declared in pdf/annot.h:47
	PDFAnnotIsToggleNoView = 256
	// PDFAnnotIsLockedContents as declared in pdf/annot.h:48
	PDFAnnotIsLockedContents = 512
)

const (
	// PDFAnnotLineEndingNone as declared in pdf/annot.h:53
	PDFAnnotLineEndingNone = iota
	// PDFAnnotLineEndingSquare as declared in pdf/annot.h:54
	PDFAnnotLineEndingSquare = 1
	// PDFAnnotLineEndingCircle as declared in pdf/annot.h:55
	PDFAnnotLineEndingCircle = 2
	// PDFAnnotLineEndingDiamond as declared in pdf/annot.h:56
	PDFAnnotLineEndingDiamond = 3
	// PDFAnnotLineEndingOpenarrow as declared in pdf/annot.h:57
	PDFAnnotLineEndingOpenarrow = 4
	// PDFAnnotLineEndingClosedarrow as declared in pdf/annot.h:58
	PDFAnnotLineEndingClosedarrow = 5
	// PDFAnnotLineEndingButt as declared in pdf/annot.h:59
	PDFAnnotLineEndingButt = 6
	// PDFAnnotLineEndingRopenarrow as declared in pdf/annot.h:60
	PDFAnnotLineEndingRopenarrow = 7
	// PDFAnnotLineEndingRclosedarrow as declared in pdf/annot.h:61
	PDFAnnotLineEndingRclosedarrow = 8
	// PDFAnnotLineEndingSlash as declared in pdf/annot.h:62
	PDFAnnotLineEndingSlash = 9
)

const ()

const (
	// PDFWidgetTypeNotWidget as declared in pdf/widget.h:7
	PDFWidgetTypeNotWidget = -1
	// PDFWidgetTypePushbutton as declared in pdf/widget.h:8
	PDFWidgetTypePushbutton = 0
	// PDFWidgetTypeCheckbox as declared in pdf/widget.h:9
	PDFWidgetTypeCheckbox = 1
	// PDFWidgetTypeRadiobutton as declared in pdf/widget.h:10
	PDFWidgetTypeRadiobutton = 2
	// PDFWidgetTypeText as declared in pdf/widget.h:11
	PDFWidgetTypeText = 3
	// PDFWidgetTypeListbox as declared in pdf/widget.h:12
	PDFWidgetTypeListbox = 4
	// PDFWidgetTypeCombobox as declared in pdf/widget.h:13
	PDFWidgetTypeCombobox = 5
	// PDFWidgetTypeSignature as declared in pdf/widget.h:14
	PDFWidgetTypeSignature = 6
)

const (
	// PDFWidgetContentUnrestrained as declared in pdf/widget.h:20
	PDFWidgetContentUnrestrained = iota
	// PDFWidgetContentNumber as declared in pdf/widget.h:21
	PDFWidgetContentNumber = 1
	// PDFWidgetContentSpecial as declared in pdf/widget.h:22
	PDFWidgetContentSpecial = 2
	// PDFWidgetContentDate as declared in pdf/widget.h:23
	PDFWidgetContentDate = 3
	// PDFWidgetContentTime as declared in pdf/widget.h:24
	PDFWidgetContentTime = 4
)

const ()

const (
	// PDFEventTypePointer as declared in pdf/event.h:13
	PDFEventTypePointer = iota
)

const (
	// PDFPointerDown as declared in pdf/event.h:19
	PDFPointerDown = iota
	// PDFPointerUp as declared in pdf/event.h:20
	PDFPointerUp = 1
)

const (
	// PDFDocumentEventAlert as declared in pdf/event.h:69
	PDFDocumentEventAlert = iota
	// PDFDocumentEventPrint as declared in pdf/event.h:70
	PDFDocumentEventPrint = 1
	// PDFDocumentEventLaunchUrl as declared in pdf/event.h:71
	PDFDocumentEventLaunchUrl = 2
	// PDFDocumentEventMailDoc as declared in pdf/event.h:72
	PDFDocumentEventMailDoc = 3
	// PDFDocumentEventSubmit as declared in pdf/event.h:73
	PDFDocumentEventSubmit = 4
	// PDFDocumentEventExecMenuItem as declared in pdf/event.h:74
	PDFDocumentEventExecMenuItem = 5
	// PDFDocumentEventExecDialog as declared in pdf/event.h:75
	PDFDocumentEventExecDialog = 6
)

const (
	// PDFAlertIconError as declared in pdf/event.h:113
	PDFAlertIconError = iota
	// PDFAlertIconWarning as declared in pdf/event.h:114
	PDFAlertIconWarning = 1
	// PDFAlertIconQuestion as declared in pdf/event.h:115
	PDFAlertIconQuestion = 2
	// PDFAlertIconStatus as declared in pdf/event.h:116
	PDFAlertIconStatus = 3
)

const (
	// PDFAlertButtonGroupOk as declared in pdf/event.h:122
	PDFAlertButtonGroupOk = iota
	// PDFAlertButtonGroupOkCancel as declared in pdf/event.h:123
	PDFAlertButtonGroupOkCancel = 1
	// PDFAlertButtonGroupYesNo as declared in pdf/event.h:124
	PDFAlertButtonGroupYesNo = 2
	// PDFAlertButtonGroupYesNoCancel as declared in pdf/event.h:125
	PDFAlertButtonGroupYesNoCancel = 3
)

const (
	// PDFAlertButtonNone as declared in pdf/event.h:131
	PDFAlertButtonNone = iota
	// PDFAlertButtonOk as declared in pdf/event.h:132
	PDFAlertButtonOk = 1
	// PDFAlertButtonCancel as declared in pdf/event.h:133
	PDFAlertButtonCancel = 2
	// PDFAlertButtonNo as declared in pdf/event.h:134
	PDFAlertButtonNo = 3
	// PDFAlertButtonYes as declared in pdf/event.h:135
	PDFAlertButtonYes = 4
)
