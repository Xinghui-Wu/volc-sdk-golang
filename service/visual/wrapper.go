package visual

import (
	"encoding/json"
	"net/url"

	"github.com/volcengine/volc-sdk-golang/service/visual/model"
)

func (p *Visual) commonHandler(api string, form url.Values, resp interface{}) (int, error) {
	respBody, statusCode, err := p.Client.Post(api, nil, form)
	if err != nil {
		errMsg := err.Error()
		// business error will be shown in resp, request error should be nil here
		if errMsg[:3] != "api" {
			return statusCode, err
		}
	}

	if err := json.Unmarshal(respBody, resp); err != nil {
		return statusCode, err
	}
	return statusCode, nil
}

func (p *Visual) BankCard(form url.Values) (*model.BankCardResult, int, error) {
	resp := new(model.BankCardResult)
	statusCode, err := p.commonHandler("BankCard", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) IDCard(form url.Values) (*model.IDCardResult, int, error) {
	resp := new(model.IDCardResult)
	statusCode, err := p.commonHandler("IDCard", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) OCRNormal(form url.Values) (*model.OCRNormalResult, int, error) {
	resp := new(model.OCRNormalResult)
	statusCode, err := p.commonHandler("OCRNormal", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) FaceSwap(form url.Values) (*model.FaceSwapResult, int, error) {
	resp := new(model.FaceSwapResult)
	statusCode, err := p.commonHandler("FaceSwap", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) JPCartoon(form url.Values) (*model.JPCartoonResult, int, error) {
	resp := new(model.JPCartoonResult)
	statusCode, err := p.commonHandler("JPCartoon", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) JPCartoonCut(form url.Values) (*model.JPCartoonCutResult, int, error) {
	resp := new(model.JPCartoonCutResult)
	statusCode, err := p.commonHandler("JPCartoonCut", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) VideoSceneDetect(form url.Values) (*model.VideoSceneDetectResult, int, error) {
	resp := new(model.VideoSceneDetectResult)
	statusCode, err := p.commonHandler("VideoSceneDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) OverResolution(form url.Values) (*model.OverResolutionResult, int, error) {
	resp := new(model.OverResolutionResult)
	statusCode, err := p.commonHandler("VideoSceneDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) GoodsSegment(form url.Values) (*model.GoodsSegmentResult, int, error) {
	resp := new(model.GoodsSegmentResult)
	statusCode, err := p.commonHandler("VideoSceneDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageOutpaint(form url.Values) (*model.ImageOutpaintResult, int, error) {
	resp := new(model.ImageOutpaintResult)
	statusCode, err := p.commonHandler("VideoSceneDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageInpaint(form url.Values) (*model.ImageInpaintResult, int, error) {
	resp := new(model.ImageInpaintResult)
	statusCode, err := p.commonHandler("VideoSceneDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) ImageCut(form url.Values) (*model.ImageCutResult, int, error) {
	resp := new(model.ImageCutResult)
	statusCode, err := p.commonHandler("VideoSceneDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) EntityDetect(form url.Values) (*model.EntityDetectResult, int, error) {
	resp := new(model.EntityDetectResult)
	statusCode, err := p.commonHandler("VideoSceneDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}

func (p *Visual) GoodsDetect(form url.Values) (*model.GoodsDetectResult, int, error) {
	resp := new(model.GoodsDetectResult)
	statusCode, err := p.commonHandler("VideoSceneDetect", form, resp)
	if err != nil {
		return nil, statusCode, err
	}
	return resp, statusCode, nil
}
