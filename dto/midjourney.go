// Copyright (c) 2025 Tethys Plex
//
// This file is part of Veloera.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>.
package dto

//type SimpleMjRequest struct {
//	Prompt   string `json:"prompt"`
//	CustomId string `json:"customId"`
//	Action   string `json:"action"`
//	Content  string `json:"content"`
//}

type SwapFaceRequest struct {
	SourceBase64 string `json:"sourceBase64"`
	TargetBase64 string `json:"targetBase64"`
}

type MidjourneyRequest struct {
	Prompt      string   `json:"prompt"`
	CustomId    string   `json:"customId"`
	BotType     string   `json:"botType"`
	NotifyHook  string   `json:"notifyHook"`
	Action      string   `json:"action"`
	Index       int      `json:"index"`
	State       string   `json:"state"`
	TaskId      string   `json:"taskId"`
	Base64Array []string `json:"base64Array"`
	Content     string   `json:"content"`
	MaskBase64  string   `json:"maskBase64"`
}

type MidjourneyResponse struct {
	Code        int         `json:"code"`
	Description string      `json:"description"`
	Properties  interface{} `json:"properties"`
	Result      string      `json:"result"`
}

type MidjourneyUploadResponse struct {
	Code        int      `json:"code"`
	Description string   `json:"description"`
	Result      []string `json:"result"`
}

type MidjourneyResponseWithStatusCode struct {
	StatusCode int `json:"statusCode"`
	Response   MidjourneyResponse
}

type MidjourneyDto struct {
	MjId        string      `json:"id"`
	Action      string      `json:"action"`
	CustomId    string      `json:"customId"`
	BotType     string      `json:"botType"`
	Prompt      string      `json:"prompt"`
	PromptEn    string      `json:"promptEn"`
	Description string      `json:"description"`
	State       string      `json:"state"`
	SubmitTime  int64       `json:"submitTime"`
	StartTime   int64       `json:"startTime"`
	FinishTime  int64       `json:"finishTime"`
	ImageUrl    string      `json:"imageUrl"`
	Status      string      `json:"status"`
	Progress    string      `json:"progress"`
	FailReason  string      `json:"failReason"`
	Buttons     any         `json:"buttons"`
	MaskBase64  string      `json:"maskBase64"`
	Properties  *Properties `json:"properties"`
}

type MidjourneyStatus struct {
	Status int `json:"status"`
}
type MidjourneyWithoutStatus struct {
	Id          int    `json:"id"`
	Code        int    `json:"code"`
	UserId      int    `json:"user_id" gorm:"index"`
	Action      string `json:"action"`
	MjId        string `json:"mj_id" gorm:"index"`
	Prompt      string `json:"prompt"`
	PromptEn    string `json:"prompt_en"`
	Description string `json:"description"`
	State       string `json:"state"`
	SubmitTime  int64  `json:"submit_time"`
	StartTime   int64  `json:"start_time"`
	FinishTime  int64  `json:"finish_time"`
	ImageUrl    string `json:"image_url"`
	Progress    string `json:"progress"`
	FailReason  string `json:"fail_reason"`
	ChannelId   int    `json:"channel_id"`
}

type ActionButton struct {
	CustomId any `json:"customId"`
	Emoji    any `json:"emoji"`
	Label    any `json:"label"`
	Type     any `json:"type"`
	Style    any `json:"style"`
}

type Properties struct {
	FinalPrompt   string `json:"finalPrompt"`
	FinalZhPrompt string `json:"finalZhPrompt"`
}
